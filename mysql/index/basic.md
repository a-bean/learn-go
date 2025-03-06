# 索引

## 1. 索引介绍

**索引是一种用于快速查询和检索数据的数据结构，其本质可以看成是一种排序好的数据结构。**

索引的作用就相当于书的目录。打个比方: 我们在查字典的时候，如果没有目录，那我们就只能一页一页的去找我们需要查的那个字，速度很慢。如果有目录了，我们只需要先去目录里查找字的位置，然后直接翻到那一页就行了。

## 2. 索引的优缺点

**优点**：

- 使用索引可以大大加快数据的检索速度（大大减少检索的数据量）, 减少 IO 次数，这也是创建索引的最主要的原因。
- 通过创建唯一性索引，可以保证数据库表中每一行数据的唯一性。

**缺点**：

- 创建索引和维护索引需要耗费许多时间。当对表中的数据进行增删改的时候，如果数据有索引，那么索引也需要动态的修改，会降低 SQL 执行效率。
- 索引需要使用物理文件存储，也会耗费一定空间。

但是，**使用索引一定能提高查询性能吗?**

大多数情况下，索引查询都是比全表扫描要快的。但是如果数据库的数据量不大，那么使用索引也不一定能够带来很大提升。

## 3. 索引底层数据结构选型

### Hash 表

哈希表是键值对的集合，通过键(key)即可快速取出对应的值(value)，因此哈希表可以快速检索数据（接近 O（1））。

**为何能够通过 key 快速取出 value 呢？** 原因在于 **哈希算法**（也叫散列算法）。通过哈希算法，我们可以快速找到 key 对应的 index，找到了 index 也就找到了对应的 value。但是！哈希算法有个 **Hash 冲突** 问题，也就是说多个不同的 key 最后得到的 index 相同。通常情况下，我们常用的解决办法是 **链地址法**。链地址法就是将哈希冲突数据存放在链表中。

MySQL 的 InnoDB 存储引擎不直接支持常规的哈希索引，但是，InnoDB 存储引擎中存在一种特殊的“自适应哈希索引”（Adaptive Hash Index），自适应哈希索引并不是传统意义上的纯哈希索引，而是结合了 B+Tree 和哈希索引的特点，以便更好地适应实际应用中的数据访问模式和性能需求。自适应哈希索引的每个哈希桶实际上是一个小型的 B+Tree 结构。这个 B+Tree 结构可以存储多个键值对，而不仅仅是一个键。这有助于减少哈希冲突链的长度，提高了索引的效率。关于 Adaptive Hash Index 的详细介绍，可以查看 [MySQL 各种“Buffer”之 Adaptive Hash Index](https://mp.weixin.qq.com/s/ra4v1XR5pzSWc-qtGO-dBg) 这篇文章。

既然哈希表这么快，**为什么 MySQL 没有使用其作为索引的数据结构呢？** 主要是因为==Hash 索引不支持顺序和范围查询==。假如我们要对表中的数据进行排序或者进行范围查询，那 Hash 索引可就不行了。并且，每次 IO 只能取一个。

```sql
SELECT * FROM tb1 WHERE id < 500;
```

Hash 索引是根据 hash 算法来定位的，难不成还要把 1 - 499 的数据，每个都进行一次 hash 计算来定位吗?这就是 Hash 最大的缺点了。

### B 树& B+树

B 树也称 B-树,全称为 **多路平衡查找树** ，B+ 树是 B 树的一种变体。B 树和 B+树中的 B 是 `Balanced` （平衡）的意思。

目前大部分数据库系统及文件系统都采用 B-Tree 或其变种 B+Tree 作为索引结构。

**B 树& B+树两者有何异同呢？**

- B 树的所有节点既存放键(key) 也存放数据(data)，而 B+树只有叶子节点存放 key 和 data，其他内节点只存放 key。
- B 树的叶子节点都是独立的;B+树的叶子节点有一条引用链指向与它相邻的叶子节点。
- B 树的检索的过程相当于对范围内的每个节点的关键字做二分查找，可能还没有到达叶子节点，检索就结束了。而 B+树的检索效率就很稳定了，任何查找都是从根节点到叶子节点的过程，叶子节点的顺序检索很明显。
- 在 B 树中进行范围查询时，首先找到要查找的下限，然后对 B 树进行中序遍历，直到找到查找的上限；而 B+树的范围查询，只需要对链表进行遍历即可。

综上，B+树与 B 树相比，具备==更少的 IO 次数、更稳定的查询效率和更适于范围查询==这些优势。

在 MySQL 中，MyISAM 引擎和 InnoDB 引擎都是使用 B+Tree 作为索引结构，但是，两者的实现方式不太一样。（下面的内容整理自《Java 工程师修炼之道》）

> MyISAM 引擎中，B+Tree 叶节点的 data 域存放的是数据记录的地址。在索引检索的时候，首先按照 B+Tree 搜索算法搜索索引，如果指定的 Key 存在，则取出其 data 域的值，然后以 data 域的值为地址读取相应的数据记录。这被称为“**非聚簇索引（非聚集索引）**”。
>
> InnoDB 引擎中，其数据文件本身就是索引文件。相比 MyISAM，索引文件和数据文件是分离的，其表数据文件本身就是按 B+Tree 组织的一个索引结构，树的叶节点 data 域保存了完整的数据记录。这个索引的 key 是数据表的主键，因此 InnoDB 表数据文件本身就是主索引。这被称为“**聚簇索引（聚集索引）**”，而其余的索引都作为 **辅助索引** ，辅助索引的 data 域存储相应记录主键的值而不是地址，这也是和 MyISAM 不同的地方。在根据主索引搜索时，直接找到 key 所在的节点即可取出数据；在根据辅助索引查找时，则需要先取出主键的值，再走一遍主索引。 因此，在设计表的时候，不建议使用过长的字段作为主键，也不建议使用非单调的字段作为主键，这样会造成主索引频繁分裂。

## 4. 索引创建

### **1. 创建普通索引（`INDEX`）**

普通索引可以加速查询，但**允许重复值**。

```sql
CREATE INDEX idx_column_name ON table_name(column_name);
```

```sql
CREATE INDEX idx_name ON users(name);
```

**适用于 `WHERE name = 'Zheng'` 加速查询**。

### **2. 创建唯一索引（`UNIQUE INDEX`）**

唯一索引确保列中的值**不重复**，适用于唯一字段，如邮箱、用户名等。

```sql
CREATE UNIQUE INDEX idx_unique_email ON users(email);
```

**等同于 `UNIQUE` 约束**：

```sql
ALTER TABLE users ADD UNIQUE (email);
```

### **3. 创建联合索引（复合索引）**

联合索引是对**多个列**创建索引，可以优化多条件查询。

```sql
CREATE INDEX idx_name_age ON users(name, age);
```

**适用于 `WHERE name = 'Zheng' AND age = 25` 的查询**。
❗ **注意索引顺序，(name, age) 适用于 `name` 开头的查询，但不适用于 `age` 单独查询。**

### **4. 创建全文索引（`FULLTEXT`）**

用于 `TEXT` 或 `VARCHAR` 类型字段，适用于**全文搜索**。

```sql
CREATE FULLTEXT INDEX idx_fulltext_content ON articles(content);
```

**用于 `MATCH(content) AGAINST('keyword')` 进行全文搜索。**
❗ **仅支持 `InnoDB` 和 `MyISAM` 引擎。**

### **5. 在 `CREATE TABLE` 时创建索引**

在建表时直接添加索引：

```sql
CREATE TABLE users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(50),
    email VARCHAR(100) UNIQUE,  -- 自动创建唯一索引
    age INT,
    INDEX idx_name (name),  -- 普通索引
    UNIQUE INDEX idx_email (email),  -- 唯一索引
    INDEX idx_name_age (name, age)  -- 联合索引
);
```

### **6. 使用 `ALTER TABLE` 添加索引**

如果表已存在，可以用 `ALTER TABLE` 添加索引：

```sql
ALTER TABLE users ADD INDEX idx_name (name);  -- 普通索引
ALTER TABLE users ADD UNIQUE INDEX idx_email (email);  -- 唯一索引
ALTER TABLE users ADD FULLTEXT INDEX idx_content (content);  -- 全文索引
```

### 7. 删除索引

如果索引不再需要，可以删除：

```sql
DROP INDEX idx_name ON users;
ALTER TABLE users DROP INDEX idx_name;
```

### 8. 自动创建索引的情况

1. `PRIMARY KEY` 和 `UNIQUE` 约束**自动创建唯一索引**。
2. `FOREIGN KEY`（InnoDB）**自动创建普通索引**。
3. `AUTO_INCREMENT` 字段**自动创建唯一索引**（通常作为主键）。
4. `FULLTEXT` 约束会自动创建**全文索引**（`FULLTEXT INDEX`）。

## 5. 索引类型总结

**按照数据结构维度划分**：

- BTree 索引：MySQL 里默认和最常用的索引类型。只有叶子节点存储 value，非叶子节点只有指针和 key。存储引擎 MyISAM 和 InnoDB 实现 BTree 索引都是使用 B+Tree，但二者实现方式不一样（前面已经介绍了）。
- 哈希索引：类似键值对的形式，一次即可定位。
- RTree 索引：一般不会使用，仅支持 geometry 数据类型，优势在于范围查找，效率较低，通常使用搜索引擎如 ElasticSearch 代替。
- 全文索引：对文本的内容进行分词，进行搜索。目前只有 `CHAR`、`VARCHAR` ，`TEXT` 列上可以创建全文索引。一般不会使用，效率较低，通常使用搜索引擎如 ElasticSearch 代替。

**按照底层存储方式角度划分**：

- 聚簇索引（聚集索引）：索引结构和数据一起存放的索引，InnoDB 中的主键索引就属于聚簇索引。
- 非聚簇索引（非聚集索引）：索引结构和数据分开存放的索引，二级索引(辅助索引)就属于非聚簇索引。MySQL 的 MyISAM 引擎，不管主键还是非主键，使用的都是非聚簇索引。

**按照应用维度划分**：

- 主键索引：加速查询 + 列值唯一（不可以有 NULL）+ 表中只有一个。

  **InnoDB 的特殊性**：

  - InnoDB 的主键索引是 **聚簇索引（Clustered Index）**，数据存储按主键顺序排列；
  - 如果没有显式指定主键，InnoDB 会选择唯一非空索引作为主键；
  - 若找不到唯一非空索引，MySQL 会自动创建一个隐藏的 `rowid` 作为主键。

  ```sql
  ALTER TABLE students ADD PRIMARY KEY (id);
  ```

- 普通索引：仅加速查询。

  ```sql
  CREATE INDEX idx_name ON students(name);
  ```

- 唯一索引：加速查询 + 列值唯一（可以有 NULL）。

  ```sql
  CREATE UNIQUE INDEX idx_email ON students(email);
  ```

- 覆盖索引：一个索引包含（或者说覆盖）所有需要查询的字段的值。

- 联合索引：多列值组成一个索引，专门用于组合搜索，其效率大于索引合并。

  ```sql
  CREATE INDEX idx_multi ON students(name, age, city);
  ```

- 全文索引：对文本的内容进行分词，进行搜索。目前只有 `CHAR`、`VARCHAR` ，`TEXT` 列上可以创建全文索引。一般不会使用，效率较低，通常使用搜索引擎如 ElasticSearch 代替。

- 前缀索引：对文本的前几个字符创建索引，相比普通索引建立的数据更小，因为只取前几个字符。

  ```sql
  CREATE INDEX idx_students_name ON students(name(3));
  ```

MySQL 8.x 中实现的索引新特性：

- 隐藏索引：也称为不可见索引，不会被优化器使用，但是仍然需要维护，通常会软删除和灰度发布的场景中使用。主键不能设置为隐藏（包括显式设置或隐式设置）。
- 降序索引：之前的版本就支持通过 desc 来指定索引为降序，但实际上创建的仍然是常规的升序索引。直到 MySQL 8.x 版本才开始真正支持降序索引。另外，在 MySQL 8.x 版本中，不再对 GROUP BY 语句进行隐式排序。
- 函数索引：从 MySQL 8.0.13 版本开始支持在索引中使用函数或者表达式的值，也就是在索引中可以包含函数或者表达式。

## 6. 主键索引(Primary Key)

数据表的主键列使用的就是主键索引。一张数据表有只能有一个主键，并且主键不能为 null，不能重复。

在 MySQL 的 InnoDB 的表中，当没有显示地指定表的主键时，InnoDB 会自动先检查表中是否有唯一索引且不允许存在 null 值的字段，如果有，则选择该字段为默认的主键，否则 InnoDB 将会自动创建一个 6Byte 的自增主键。

![主键索引](https://oss.javaguide.cn/github/javaguide/open-source-project/cluster-index.png)

## 7. 二级索引

二级索引（Secondary Index）的叶子节点存储的数据是==主键==的值，也就是说，通过二级索引可以定位主键的位置，二级索引又称为辅助索引/非主键索引。

唯一索引，普通索引，前缀索引等索引都属于二级索引。

1. **唯一索引(Unique Key)**:唯一索引也是一种约束。唯一索引的属性列不能出现重复的数据，但是允许数据为 NULL，一张表允许创建多个唯一索引。 建立唯一索引的目的大部分时候都是为了该属性列的数据的唯一性，而不是为了查询效率。
2. **普通索引(Index)**:普通索引的唯一作用就是为了快速查询数据，一张表允许创建多个普通索引，并允许数据重复和 NULL。
3. **前缀索引(Prefix)**:前缀索引只适用于字符串类型的数据。前缀索引是对文本的前几个字符创建索引，相比普通索引建立的数据更小，因为只取前几个字符。
4. **全文索引(Full Text)**:全文索引主要是为了检索大文本数据中的关键字的信息，是目前搜索引擎数据库使用的一种技术。Mysql5.6 之前只有 MYISAM 引擎支持全文索引，5.6 之后 InnoDB 也支持了全文索引。

二级索引:

![二级索引](https://oss.javaguide.cn/github/javaguide/open-source-project/no-cluster-index.png)

### 聚簇索引（聚集索引）

#### 聚簇索引介绍

聚簇索引（Clustered Index）即索引结构和数据一起存放的索引，并不是一种单独的索引类型。InnoDB 中的主键索引就属于聚簇索引。

在 MySQL 中，InnoDB 引擎的表的 `.ibd`文件就包含了该表的索引和数据，对于 InnoDB 引擎表来说，该表的索引(B+树)的每个非叶子节点存储索引，叶子节点存储索引和索引对应的数据。

#### 聚簇索引的优缺点

**优点**：

- **查询速度非常快**：聚簇索引的查询速度非常的快，因为整个 B+树本身就是一颗多叉平衡树，叶子节点也都是有序的，定位到索引的节点，就相当于定位到了数据。相比于非聚簇索引， 聚簇索引少了一次读取数据的 IO 操作。
- **对排序查找和范围查找优化**：聚簇索引对于主键的排序查找和范围查找速度非常快。

**缺点**：

- **依赖于有序的数据**：因为 B+树是多路平衡树，如果索引的数据不是有序的，那么就需要在插入时排序，如果数据是整型还好，否则类似于字符串或 UUID 这种又长又难比较的数据，插入或查找的速度肯定比较慢。
- **更新代价大**：如果对索引列的数据被修改时，那么对应的索引也将会被修改，而且聚簇索引的叶子节点还存放着数据，修改代价肯定是较大的，所以对于主键索引来说，主键一般都是不可被修改的。

### 非聚簇索引（非聚集索引）

#### 非聚簇索引介绍

非聚簇索引(Non-Clustered Index)即索引结构和数据分开存放的索引，并不是一种单独的索引类型。二级索引(辅助索引)就属于非聚簇索引。MySQL 的 MyISAM 引擎，不管主键还是非主键，使用的都是非聚簇索引。

非聚簇索引的叶子节点并不一定存放数据的指针，因为二级索引的叶子节点就存放的是主键，根据主键再回表查数据。

#### 非聚簇索引的优缺点

**优点**：更新代价比聚簇索引要小 。非聚簇索引的更新代价就没有聚簇索引那么大了，非聚簇索引的叶子节点是不存放数据的。

**缺点**：

- **依赖于有序的数据**:跟聚簇索引一样，非聚簇索引也依赖于有序的数据
- **可能会二次查询(回表)**:这应该是非聚簇索引最大的缺点了。 当查到索引对应的指针或主键后，可能还需要根据指针或主键再到数据文件或表中查询。

这是 MySQL 的表的文件截图:

![MySQL 表的文件](https://oss.javaguide.cn/github/javaguide/database/mysql20210420165311654.png)

聚簇索引和非聚簇索引:

![聚簇索引和非聚簇索引](https://oss.javaguide.cn/github/javaguide/database/mysql20210420165326946.png)

#### 非聚簇索引一定回表查询吗(覆盖索引)?

**非聚簇索引不一定回表查询。**

试想一种情况，用户准备使用 SQL 查询用户名，而用户名字段正好建立了索引。

```sql
 SELECT name FROM table WHERE name='guang19';
```

那么这个索引的 key 本身就是 name，查到对应的 name 直接返回就行了，无需==回表查询==。

即使是 MYISAM 也是这样，虽然 MYISAM 的主键索引确实需要回表，因为它的主键索引的叶子节点存放的是指针。但是！**如果 SQL 查的就是主键呢?**

```sql
SELECT id FROM table WHERE id=1;
```

主键索引本身的 key 就是主键，查到返回就行了。这种情况就称之为覆盖索引了。

## 8. 覆盖索引和联合索引

### 覆盖索引

如果一个索引包含（或者说覆盖）所有需要查询的字段的值，我们就称之为 **覆盖索引（Covering Index）** 。

在 InnoDB 存储引擎中，非主键索引的叶子节点包含的是主键的值。这意味着，当使用非主键索引进行查询时，数据库会先找到对应的主键值，然后再通过主键索引来定位和检索完整的行数据。这个过程被称为“回表”。

**覆盖索引即需要查询的字段正好是索引的字段，那么直接根据该索引，就可以查到数据了，而无需回表查询。**

> 如主键索引，如果一条 SQL 需要查询主键，那么正好根据主键索引就可以查到主键。再如普通索引，如果一条 SQL 需要查询 name，name 字段正好有索引，
> 那么直接根据这个索引就可以查到数据，也无需回表。

![覆盖索引](https://oss.javaguide.cn/github/javaguide/database/mysql20210420165341868.png)

我们这里简单演示一下覆盖索引的效果。

1、创建一个名为 `cus_order` 的表，来实际测试一下这种排序方式。为了测试方便， `cus_order` 这张表只有 `id`、`score`、`name`这 3 个字段。

```sql
CREATE TABLE `cus_order` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `score` int(11) NOT NULL,
  `name` varchar(11) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4;
```

2、定义一个简单的存储过程（PROCEDURE）来插入 100w 测试数据。

```sql
DELIMITER ;;
CREATE DEFINER=`root`@`%` PROCEDURE `BatchinsertDataToCusOder`(IN start_num INT,IN max_num INT)
BEGIN
      DECLARE i INT default start_num;
      WHILE i < max_num DO
          insert into `cus_order`(`id`, `score`, `name`)
          values (i,RAND() * 1000000,CONCAT('user', i));
          SET i = i + 1;
      END WHILE;
  END;;
DELIMITER ;
```

存储过程定义完成之后，我们执行存储过程即可！

```sql
CALL BatchinsertDataToCusOder(1, 1000000); # 插入100w+的随机数据
```

等待一会，100w 的测试数据就插入完成了！

3、创建覆盖索引并使用 `EXPLAIN` 命令分析。

为了能够对这 100w 数据按照 `score` 进行排序，我们需要执行下面的 SQL 语句。

```sql
#降序排序
SELECT `score`,`name` FROM `cus_order` ORDER BY `score` DESC;
```

使用 `EXPLAIN` 命令分析这条 SQL 语句，通过 `Extra` 这一列的 `Using filesort` ，我们发现是没有用到覆盖索引的。

![img](https://oss.javaguide.cn/github/javaguide/mysql/not-using-covering-index-demo.png)

不过这也是理所应当，毕竟我们现在还没有创建索引呢！

我们这里以 `score` 和 `name` 两个字段建立联合索引：

```sql
ALTER TABLE `cus_order` ADD INDEX id_score_name(score, name);
```

创建完成之后，再用 `EXPLAIN` 命令分析再次分析这条 SQL 语句。

![img](https://oss.javaguide.cn/github/javaguide/mysql/using-covering-index-demo.png)

通过 `Extra` 这一列的 `Using index` ，说明这条 SQL 语句成功使用了覆盖索引。

### 联合索引

使用表中的多个字段创建索引，就是 **联合索引**，也叫 **组合索引** 或 **复合索引**。

以 `score` 和 `name` 两个字段建立联合索引：

```sql
ALTER TABLE `cus_order` ADD INDEX id_score_name(score, name);
```

### 最左前缀匹配原则

最左前缀匹配原则指的是在使用联合索引时，MySQL 会根据索引中的字段顺序，从左到右依次匹配查询条件中的字段。如果查询条件与索引中的最左侧字段相匹配，那么 MySQL 就会使用索引来过滤数据，这样可以提高查询效率。

最左匹配原则会一直向右匹配，直到遇到范围查询（如 >、<）为止。对于 >=、<=、BETWEEN 以及前缀匹配 LIKE 的范围查询，不会停止匹配（相关阅读：[联合索引的最左匹配原则全网都在说的一个错误结论](https://mp.weixin.qq.com/s/8qemhRg5MgXs1So5YCv0fQ)）。

假设有一个联合索引`(column1, column2, column3)`，其从左到右的所有前缀为`(column1)`、`(column1, column2)`、`(column1, column2, column3)`（创建 1 个联合索引相当于创建了 3 个索引），包含这些列的所有查询都会走索引而不会全表扫描。

我们在使用联合索引时，可以将区分度高的字段放在最左边，这也可以过滤更多数据。

我们这里简单演示一下最左前缀匹配的效果。

1、创建一个名为 `student` 的表，这张表只有 `id`、`name`、`class`这 3 个字段。

```sql
CREATE TABLE `student` (
  `id` int NOT NULL,
  `name` varchar(100) DEFAULT NULL,
  `class` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`id`),
  INDEX `name_class_idx` (`name`,`class`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
```

2、下面我们分别测试三条不同的 SQL 语句。

![img](https://oss.javaguide.cn/github/javaguide/database/mysql/leftmost-prefix-matching-rule.png)

```sql
# 可以命中索引
SELECT * FROM student WHERE name = 'Anne Henry';
EXPLAIN SELECT * FROM student WHERE name = 'Anne Henry' AND class = 'lIrm08RYVk';
# 无法命中索引
SELECT * FROM student WHERE class = 'lIrm08RYVk';
```

再来看一个常见的面试题：如果有索引 `联合索引（a，b，c）`，查询 `a=1 AND c=1`会走索引么？`c=1` 呢？`b=1 AND c=1`呢？

先不要往下看答案，给自己 3 分钟时间想一想。

1. 查询 `a=1 AND c=1`：根据最左前缀匹配原则，查询可以使用索引的前缀部分。因此，该查询仅在 `a=1` 上使用索引，然后对结果进行 `c=1` 的过滤。
2. 查询 `c=1` ：由于查询中不包含最左列 `a`，根据最左前缀匹配原则，整个索引都无法被使用。
3. 查询`b=1 AND c=1`：和第二种一样的情况，整个索引都不会使用。

MySQL 8.0.13 版本引入了索引跳跃扫描（Index Skip Scan，简称 ISS），它可以在某些索引查询场景下提高查询效率。在没有 ISS 之前，不满足最左前缀匹配原则的联合索引查询中会执行全表扫描。而 ISS 允许 MySQL 在某些情况下避免全表扫描，即使查询条件不符合最左前缀。不过，这个功能比较鸡肋， 和 Oracle 中的没法比，MySQL 8.0.31 还报告了一个 bug：[Bug #109145 Using index for skip scan cause incorrect result](https://bugs.mysql.com/bug.php?id=109145)（后续版本已经修复）。个人建议知道有这个东西就好，不需要深究，实际项目也不一定能用上。

## 9. 索引下推

**索引下推（Index Condition Pushdown，简称 ICP）** 是 **MySQL 5.6** 版本中提供的一项索引优化功能，==它允许存储引擎在索引遍历过程中，执行部分 `WHERE`字句的判断条件，直接过滤掉不满足条件的记录，从而减少回表次数，提高查询效率==。

假设我们有一个名为 `user` 的表，其中包含 `id`, `username`, `zipcode`和 `birthdate` 4 个字段，创建了联合索引`(zipcode, birthdate)`。

```sql
CREATE TABLE `user` (
  `id` int NOT NULL AUTO_INCREMENT,
  `username` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `zipcode` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `birthdate` date NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `idx_username_birthdate` (`zipcode`,`birthdate`)
) ENGINE=InnoDB AUTO_INCREMENT=1001 DEFAULT CHARSET=utf8mb4;

# 查询 zipcode 为 431200 且生日在 3 月的用户
# birthdate 字段使用函数索引失效
SELECT * FROM user WHERE zipcode = '431200' AND MONTH(birthdate) = 3;
```

- 没有索引下推之前，即使 `zipcode` 字段利用索引可以帮助我们快速定位到 `zipcode = '431200'` 的用户，但我们仍然需要对每一个找到的用户进行回表操作，获取完整的用户数据，再去判断 `MONTH(birthdate) = 3`。
- 有了索引下推之后，存储引擎会在使用`zipcode` 字段索引查找`zipcode = '431200'` 的用户时，同时判断`MONTH(birthdate) = 3`。这样，只有同时满足条件的记录才会被返回，减少了回表次数。

![img](https://oss.javaguide.cn/github/javaguide/database/mysql/index-condition-pushdown.png)

![img](https://oss.javaguide.cn/github/javaguide/database/mysql/index-condition-pushdown-graphic-illustration.png)

再来讲讲索引下推的具体原理，先看下面这张 MySQL 简要架构图。

![img](https://oss.javaguide.cn/javaguide/13526879-3037b144ed09eb88.png)

MySQL 可以简单分为 Server 层和存储引擎层这两层。Server 层处理查询解析、分析、优化、缓存以及与客户端的交互等操作，而存储引擎层负责数据的存储和读取，MySQL 支持 InnoDB、MyISAM、Memory 等多种存储引擎。

索引下推的**下推**其实就是指将部分上层（Server 层）负责的事情，交给了下层（存储引擎层）去处理。

我们这里结合索引下推原理再对上面提到的例子进行解释。

没有索引下推之前：

- 存储引擎层先根据 `zipcode` 索引字段找到所有 `zipcode = '431200'` 的用户的主键 ID，然后二次回表查询，获取完整的用户数据；
- 存储引擎层把所有 `zipcode = '431200'` 的用户数据全部交给 Server 层，Server 层根据`MONTH(birthdate) = 3`这一条件再进一步做筛选。

有了索引下推之后：

- 存储引擎层先根据 `zipcode` 索引字段找到所有 `zipcode = '431200'` 的用户，然后直接判断 `MONTH(birthdate) = 3`，筛选出符合条件的主键 ID；
- 二次回表查询，根据符合条件的主键 ID 去获取完整的用户数据；
- 存储引擎层把符合条件的用户数据全部交给 Server 层。

可以看出，**除了可以减少回表次数之外，索引下推还可以减少存储引擎层和 Server 层的数据传输量。**

最后，总结一下索引下推应用范围：

1. 适用于 InnoDB 引擎和 MyISAM 引擎的查询。
2. 适用于执行计划是 range, ref, eq_ref, ref_or_null 的范围查询。
3. 对于 InnoDB 表，仅用于非聚簇索引。索引下推的目标是减少全行读取次数，从而减少 I/O 操作。对于 InnoDB 聚集索引，完整的记录已经读入 InnoDB 缓冲区。在这种情况下使用索引下推 不会减少 I/O。
4. 子查询不能使用索引下推，因为子查询通常会创建临时表来处理结果，而这些临时表是没有索引的。
5. 存储过程不能使用索引下推，因为存储引擎无法调用存储函数。
