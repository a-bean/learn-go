# **MySQL 性能分析详解**

MySQL 性能分析涉及 SQL 执行计划、索引优化、查询日志分析、数据库参数调整等方面。下面详细介绍 MySQL 性能分析的关键技术、工具和优化方法。

## **一、MySQL 性能分析的关键指标**

在分析 MySQL 性能时，我们主要关注以下指标：

1. **查询执行时间**：SQL 语句的运行时间，单位 `ms`。
2. **CPU 使用率**：MySQL 进程占用的 CPU 资源情况。
3. **I/O 负载**：磁盘和网络 I/O 操作，影响查询性能。
4. 查询吞吐量（QPS/TPS）：
   - **QPS（Queries Per Second）**：每秒查询次数。
   - **TPS（Transactions Per Second）**：每秒事务处理数。
5. **连接数**：当前并发连接数，受 `max_connections` 限制。
6. **锁等待**：锁竞争情况，影响事务执行速度。
7. 缓存命中率：
   - **InnoDB Buffer Pool 命中率**（越高越好）。
   - **查询缓存（Query Cache）命中率**（MySQL 8.0 已移除）。
8. **慢查询**：耗时较长的 SQL 语句分析。
9. **索引使用情况**：是否有效利用索引，避免全表扫描。

## **二、MySQL 性能分析工具**

### **1. EXPLAIN（SQL 执行计划分析）**

`EXPLAIN` 是 MySQL 提供的 SQL 语句分析工具，可以显示 SQL 语句的执行计划，帮助优化查询性能。可以分析 SQL 语句的执行计划，判断是否使用了索引。

```sql
EXPLAIN SELECT * FROM students WHERE name = '张三';
-- 或者
EXPLAIN FORMAT=JSON SELECT * FROM students WHERE name = '张三';
```

执行 `EXPLAIN` 后，返回的结果包括多个重要字段，如下：

| 字段            | 含义             | 影响查询优化                                                                                                                                                                                                                                            |
| --------------- | ---------------- | :------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `id`            | 查询的唯一标识符 | `id` 值相同时，从上到下执行。`id` 值不同，值越大优先执行。`id` 差值较大时，表示子查询。                                                                                                                                                                 |
| `select_type`   | 查询类型         | SIMPLE（简单查询（无 `JOIN`、子查询））、PRIMARY（最外层查询）、SUBQUERY（子查询）、DERIVED（派生表）、 UNION（`UNION` 语句的第二个及之后的查询）、UNION RESULT（`UNION` 结果存放的临时表）                                                             |
| `table`         | 查询涉及的表     | 表名                                                                                                                                                                                                                                                    |
| `partitions`    | 使用的分区       | 仅适用于分区表                                                                                                                                                                                                                                          |
| `type`          | 访问类型         | 1. system：表只有 1 行数据（最优）、2. const：主键/唯一索引等值查询（最优）、3. eq_ref：主键/唯一索引关联查询（较优）、4. ref：普通索引等值查询、5. range：索引范围查询、6. index：全索引扫描（不如 `range`）7. ALL：全表扫描（最差）【越靠后性能越差】 |
| `possible_keys` | 可能使用的索引   | 但不一定真的用                                                                                                                                                                                                                                          |
| `key`           | 实际使用的索引   | 若为 NULL，则未使用索引                                                                                                                                                                                                                                 |
| `key_len`       | 索引字段长度     | 越短越好                                                                                                                                                                                                                                                |
| `ref`           | 索引比较的列     | 主要用于 `ref` 访问类型                                                                                                                                                                                                                                 |
| `rows`          | 预计扫描的行数   | 数值越小越好                                                                                                                                                                                                                                            |
| `filtered`      | 结果过滤百分比   | 100% 表示所有行都符合条件                                                                                                                                                                                                                               |
| `extra`         | 额外信息         | 是否使用临时表、排序等                                                                                                                                                                                                                                  |

**优化目标**：让 `type` 尽量接近 `const`、`ref`，避免 `ALL`（全表扫描）。

### **2. SHOW PROFILE（SQL 详细执行过程分析）**

`SHOW PROFILE` 是 MySQL 提供的 SQL 语句性能分析工具，可以查看 SQL 语句在执行过程中各个阶段的耗时，从而找出性能瓶颈，优化查询。

**注意**：`SHOW PROFILE` 在 MySQL 5.0.37 引入，但 MySQL 8.0 及以上已移除，推荐使用 `performance_schema` 代替。

```sql
-- 默认 profiling 是关闭的，需手动开启：
SET profiling = 1;
SELECT * FROM students WHERE name = '张三';
SHOW PROFILES;
```

然后查看具体的 SQL 执行步骤：

```sql
SHOW PROFILE FOR QUERY 1;
```

> **返回结果示例** | Status | Duration | |-------------------|----------| | Starting | 0.0001 | | Checking permissions | 0.0002 | | Opening tables | 0.0003 | | System lock | 0.0001 | | Table lock | 0.0002 | | Optimizing | 0.0004 | | Statistics | 0.0005 | | Preparing | 0.0006 | | Executing | 0.0007 | | Sending data | 0.0015 | | End | 0.0001 | | Query end | 0.0001 | | Closing tables | 0.0002 | | Freeing items | 0.0001 | | Logging slow query | 0.0001 | | Cleaning up | 0.0002 |

#### **优化策略**

| 问题                    | 解决方案                   |
| ----------------------- | -------------------------- |
| `Sending data` 耗时长   | 索引优化，减少扫描行数     |
| `Sorting result` 耗时长 | `ORDER BY` 字段加索引      |
| `Creating tmp table`    | 避免 `GROUP BY`，增加索引  |
| `Opening tables` 慢     | 可能是表锁竞争，可优化事务 |

#### 替代方案: MySQL 8.0 及以上

由于 `SHOW PROFILE` 在 MySQL 8.0 被移除，推荐使用 `performance_schema`。

1. 启用 `performance_schema`

```sql
UPDATE performance_schema.setup_instruments SET ENABLED = 'YES', TIMED = 'YES';
```

2. 查询 SQL 执行时间

```sql
SELECT * FROM performance_schema.events_statements_summary_by_digest ORDER BY AVG_TIMER_WAIT DESC LIMIT 5;
```

> 该查询返回 **执行最慢的 SQL 语句**。

3. 查看等待事件

```sql
SELECT EVENT_NAME, SUM_TIMER_WAIT/1000000000 AS wait_time_ms
FROM performance_schema.events_waits_summary_global_by_event_name
ORDER BY wait_time_ms DESC
LIMIT 10;
```

> 该查询用于 **分析 SQL 语句的等待瓶颈**（锁、磁盘 IO 等）。

4. SHOW STATUS（MySQL 运行状态监控）\*\*

`SHOW STATUS` 监控 MySQL 服务器状态，常用的指标如下：

**常用查询**

```sql
SHOW GLOBAL STATUS LIKE 'Threads%'; -- 线程相关
SHOW GLOBAL STATUS LIKE 'Connections'; -- 连接数
SHOW GLOBAL STATUS LIKE 'Slow_queries'; -- 慢查询数
SHOW GLOBAL STATUS LIKE 'Innodb_buffer_pool_read_requests'; -- 缓存命中率
```

**关键指标**

| 变量                             | 含义                 |
| -------------------------------- | -------------------- |
| Threads_connected                | 当前连接的线程数     |
| Threads_running                  | 正在执行的线程数     |
| Connections                      | 连接总数             |
| Slow_queries                     | 慢查询的次数         |
| Innodb_buffer_pool_read_requests | InnoDB 缓存命中率    |
| Innodb_buffer_pool_reads         | 未命中缓存的读操作数 |

> **优化目标**：调整 MySQL 配置参数，提高缓冲池命中率，减少慢查询。

### 3. SHOW STATUS（MySQL 运行状态监控）

`SHOW STATUS` 是 MySQL 提供的 **数据库运行状态** 查询工具，作用如下：

- 监控 **MySQL 服务器运行状态**，如连接数、查询数、缓存使用情况等；
- 分析 **SQL 查询性能**，如慢查询、索引命中率等；
- 诊断 **系统瓶颈**，如锁等待、线程消耗、事务状态等。

#### 1. 语法

```sql
SHOW [GLOBAL | SESSION] STATUS [LIKE 'pattern'];
```

- **`GLOBAL`**：显示 MySQL **整个服务器** 的状态；
- **`SESSION`**：显示 **当前会话** 的状态（默认）；
- **`LIKE 'pattern'`**：用于匹配特定的状态变量（支持通配符 `%`）。

#### 2. SHOW STATUS 关键指标解析

**连接数相关**:

| 变量                   | 说明             | 作用                       |
| ---------------------- | ---------------- | -------------------------- |
| `Connections`          | 总连接请求次数   | 反映数据库连接频率         |
| `Threads_connected`    | 当前连接数       | 过高可能导致资源占用问题   |
| `Threads_running`      | 正在运行的线程数 | 过高可能是 SQL 性能瓶颈    |
| `Max_used_connections` | 历史最大连接数   | 用于调整 `max_connections` |

**优化建议**：

- `Threads_connected` 过高，可调整 `max_connections`；
- `Connections` 过高，可能应用未正确使用连接池（如 MySQL 连接泄漏）。

#### **3. 查询性能相关**

| 变量           | 说明                  | 作用                               |
| -------------- | --------------------- | ---------------------------------- |
| `Queries`      | 总查询数              | 评估数据库负载                     |
| `Com_select`   | `SELECT` 语句执行次数 | 反映查询频率                       |
| `Com_insert`   | `INSERT` 语句执行次数 | 监控写入频率                       |
| `Com_update`   | `UPDATE` 语句执行次数 | 监控更新频率                       |
| `Com_delete`   | `DELETE` 语句执行次数 | 监控删除频率                       |
| `Slow_queries` | 慢查询次数            | 可结合 `slow_query_log` 进一步分析 |

**优化建议**：

- 如果 `Slow_queries`过高，启用 `slow_query_log`并使用 `EXPLAIN` 进行优化：

  ```sql
  SET GLOBAL slow_query_log = 1;
  SET GLOBAL long_query_time = 2; -- 记录执行时间大于 2 秒的查询
  ```

- 使用 `SHOW PROCESSLIST`查看当前运行的 SQL：

  ```sql
  SHOW FULL PROCESSLIST;
  ```

#### 4 索引和查询缓存

| 变量                | 说明             | 作用                |
| ------------------- | ---------------- | ------------------- |
| `Handler_read_rnd`  | 全表扫描次数     | 过高说明索引未命中  |
| `Handler_read_key`  | 通过索引读取次数 | 高代表索引利用率高  |
| `Handler_read_next` | 通过索引扫描次数 | 适用于 `range` 查询 |
| `Qcache_hits`       | 查询缓存命中次数 | 8.0 已废弃          |
| `Qcache_inserts`    | 插入缓存的查询数 | 8.0 已废弃          |

**优化建议**：

- `Handler_read_rnd` 过高 → **优化索引**，避免 `SELECT *`；
- `Handler_read_key` 过低 → **检查是否使用索引** (`EXPLAIN`)；
- `Qcache_hits` 低（8.0 以前）→ **考虑调整 `query_cache_size`**。

---

#### **5. 事务 & InnoDB 相关**

| 变量                               | 说明               | 作用                                |
| ---------------------------------- | ------------------ | ----------------------------------- |
| `Innodb_buffer_pool_read_requests` | 从缓冲池读取次数   | 高表示缓存命中率高                  |
| `Innodb_buffer_pool_reads`         | 直接从磁盘读取次数 | 低表示缓冲池足够大                  |
| `Innodb_row_lock_time`             | 行锁等待时间       | 过高可能导致事务瓶颈                |
| `Innodb_row_lock_waits`            | 行锁等待次数       | 过高可能是索引或事务问题            |
| `Innodb_log_waits`                 | 事务日志等待次数   | 过高需调整 `innodb_log_buffer_size` |

**优化建议**：

- `Innodb_buffer_pool_reads` 过高 → **增大 `innodb_buffer_pool_size`**；

- `Innodb_row_lock_waits`过高 → 优化事务，减少锁冲突：

  ```sql
  SET GLOBAL innodb_lock_wait_timeout = 10;
  ```

- `Innodb_log_waits` 过高 → **增大 `innodb_log_buffer_size`**。

#### **6 锁 & 并发**

| 变量                    | 说明             | 作用                      |
| ----------------------- | ---------------- | ------------------------- |
| `Table_locks_waited`    | 表级锁等待次数   | 过高说明表锁争用严重      |
| `Table_locks_immediate` | 表锁立即成功次数 | 高说明锁竞争较少          |
| `Threads_cached`        | 线程缓存命中数   | 高说明线程缓存工作良好    |
| `Threads_created`       | 创建新线程数     | 过高可能导致 CPU 资源浪费 |

**优化建议**：

- `Table_locks_waited` 过高 → **使用 InnoDB** 代替 MyISAM；

- `Threads_created`过高 → 增大 `thread_cache_size`：

  ```sql
  SET GLOBAL thread_cache_size = 16;
  ```

#### **7. 查询 MySQL 服务器状态**

**1 查询当前连接数**

```sql
SHOW STATUS LIKE 'Threads_connected';
```

**2 查询数据库查询总数**

```sql
SHOW STATUS LIKE 'Queries';
```

**3 查询慢查询次数**

```sql
SHOW STATUS LIKE 'Slow_queries';
```

**4 查询索引命中率**

```sql
SHOW STATUS WHERE Variable_name IN ('Handler_read_rnd', 'Handler_read_key');
```

- 索引命中率计算：

  ```sql
  SELECT (Handler_read_key / (Handler_read_key + Handler_read_rnd)) * 100 AS index_hit_rate;
  ```

#### 8. MySQL 8.0 替代方案

由于 `SHOW STATUS` 只能提供部分性能数据，MySQL 8.0 推荐使用 `performance_schema`：

**查询活跃连接**

```sql
SELECT * FROM performance_schema.threads WHERE PROCESSLIST_STATE IS NOT NULL;
```

**查询锁等待**

```sql
SELECT * FROM performance_schema.data_lock_waits;
```

### **4. SHOW PROCESSLIST（实时查询监控）**

`SHOW PROCESSLIST` 可以查看当前正在执行的 SQL 语句，发现慢查询或死锁。

```sql
SHOW PROCESSLIST;
```

**输出**

| ID  | 用户  | 进程  | 状态         | 执行时间 | SQL 语句                        |
| --- | ----- | ----- | ------------ | -------- | ------------------------------- |
| 123 | root  | Query | Sending data | 10s      | SELECT \* FROM orders WHERE ... |
| 124 | user1 | Sleep | 0            |          |                                 |

> **优化目标**：找出长时间运行的 SQL，优化索引或分解查询。

### **5. 慢查询日志（Slow Query Log）**

MySQL 可以记录执行时间超过 `long_query_time` 的 SQL 语句，便于分析慢查询。

**开启慢查询日志**

```sql
SET GLOBAL slow_query_log = 1;
SET GLOBAL long_query_time = 2; -- 记录执行时间超过 2 秒的查询
```

**查询慢查询日志**

```sql
SELECT * FROM mysql.slow_log ORDER BY start_time DESC LIMIT 10;
```

> **优化目标**：找到执行慢的 SQL，优化索引或改写 SQL 语句。

### **6. PERFORMANCE_SCHEMA（高级性能分析）**

`performance_schema` 是 MySQL 5.5 及以上版本内置的性能分析工具，可以监控锁、线程、等待事件等。

**查看是否启用**

```sql
SHOW VARIABLES LIKE 'performance_schema';
```

**查询等待时间最长的 SQL**

```sql
SELECT event_name, SUM_TIMER_WAIT/1000000000 AS wait_time_ms
FROM performance_schema.events_waits_summary_global_by_event_name
ORDER BY wait_time_ms DESC
LIMIT 10;
```

> **优化目标**：找出等待时间长的 SQL，优化事务或索引。

## 三、索引优化策略

### 1. 选择合适的索引

✅ **主键（`PRIMARY KEY`）**

```sql
CREATE TABLE users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(50)
);
```

✅ **唯一索引（`UNIQUE INDEX`）**

```sql
CREATE UNIQUE INDEX idx_email ON users(email);
```

✅ **普通索引（`INDEX`）**

```sql
CREATE INDEX idx_name ON users(name);
```

✅ **联合索引（`COMPOSITE INDEX`）**

```sql
CREATE INDEX idx_name_age ON users(name, age);
```

✅ **全文索引（`FULLTEXT INDEX`）**

```sql
CREATE FULLTEXT INDEX idx_content ON articles(content);
```

### 2. 最左前缀匹配原则

**联合索引 `(name, age)` 只能匹配以下查询：**

```sql
SELECT * FROM users WHERE name = 'Tom';  ✅ 命中索引
SELECT * FROM users WHERE name = 'Tom' AND age = 25; ✅ 命中索引
SELECT * FROM users WHERE age = 25; ❌ 索引失效
```

**解决方案**：
如果需要 `age` 独立查询，单独创建索引：

```sql
CREATE INDEX idx_age ON users(age);
```

### 3. 索引覆盖（Covering Index）

**索引覆盖 = 查询的数据列全部在索引中，无需回表**

```sql
CREATE INDEX idx_email ON users(email);
SELECT email FROM users WHERE email = 'test@example.com';  ✅ 覆盖索引
```

**`EXPLAIN` 显示 `Using index`，表示索引覆盖生效**

### 4. 减少索引大小（前缀索引）

对于长字符串字段，如 `VARCHAR(255)`，可以使用**前缀索引**：

```sql
CREATE INDEX idx_email ON users(email(10));
```

- ✅ **节省存储空间**
- ✅ **提高查询效率**
- ❗ **可能会增加重复值，需权衡长度**

### 5. 避免索引失效

✅ **使用相同的数据类型**

```sql
SELECT * FROM users WHERE id = '100';  ❌ 索引失效
SELECT * FROM users WHERE id = 100;  ✅ 索引生效
```

✅ **避免 `OR` 导致索引失效**

```sql
SELECT * FROM users WHERE name = 'Tom' OR age = 25;  ❌ 索引失效
```

✅ **优化 `OR` 语句**

```sql
SELECT * FROM users WHERE name = 'Tom'
UNION ALL
SELECT * FROM users WHERE age = 25;
```

✅ **避免 `LIKE '%xxx%'`**

```sql
SELECT * FROM users WHERE name LIKE '%Tom%';  ❌ 索引失效
SELECT * FROM users WHERE name LIKE 'Tom%';  ✅ 索引生效
```

✅ **避免对索引列使用函数、列运算**,

```sql
SELECT * FROM users WHERE LEFT(name, 3) = 'Tom';  ❌ 索引失效
SELECT * FROM users WHERE name LIKE 'Tom%';  ✅ 索引生效
```

✅ **避免隐式类型转换**

```sql
SELECT * FROM users WHERE phone = 13800001234;  ❌ 索引失效
SELECT * FROM users WHERE phone = '13800001234';  ✅ 索引生效
```

## 四. 索引设计原则 

1. 针对于数据量较大，且查询比较频繁的表建立索引。 

2. 针对于常作为查询条件（where）、排序（order by）、分组（group by）操作的字段建立索 引。

3. 尽量选择区分度高的列作为索引，尽量建立唯一索引，区分度越高，使用索引的效率越高。 

4. 如果是字符串类型的字段，字段的长度较长，可以针对于字段的特点，建立前缀索引。

5. 尽量使用联合索引，减少单列索引，查询时，联合索引很多时候可以覆盖索引，节省存储空间， 避免回表，提高查询效率。 

6. 要控制索引的数量，索引并不是多多益善，索引越多，维护索引结构的代价也就越大，会影响增 删改的效率。 1 create unique index idx_user_phone_name on tb_user(phone,name); 

7. 如果索引列不能存储NULL值，请在创建表时使用NOT NULL约束它。当优化器知道每列是否包含 NULL值时，它可以更好地确定哪个索引最有效地用于查询。
