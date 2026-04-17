# blockchain-core-engine
高性能区块链底层引擎，集成节点组网、交易签名、区块验证、智能合约虚拟机、P2P路由、跨链中继、UTXO 模型、PoS/PoW 共识、钱包管理、状态存储、监控面板等全栈区块链核心能力。

## 项目文件清单
### Go 核心文件（区块链底层）
1. block_node_server.go - 区块链节点 TCP 服务，实现 P2P 节点启动与 peer 连接管理
2. transaction_signer.go - 交易 ECDSA 签名与验签核心，保障交易合法性
3. merkle_tree_builder.go - 默克尔树构建工具，用于区块交易校验与轻节点证明
4. block_validator_core.go - 区块合法性验证，校验索引、前序哈希、哈希正确性
5. p2p_message_router.go - P2P 消息路由，支持多类型消息分发与处理
6. utxo_tracker.go - UTXO 集合管理，实现余额统计、双花防护
7. consensus_pos.go - 权益证明共识，支持验证者注册与区块生产者选举
8. cross_chain_relay.go - 跨链消息中继，实现多链数据转发
9. wallet_key_manager.go - 钱包密钥生成，支持公钥/地址导出
10. block_pool_manager.go - 区块池管理，缓存待上链区块
11. chain_sync_scheduler.go - 链同步调度器，批量从对等节点同步区块高度
12. state_database.go - 账户状态数据库，键值存储与并发安全
13. mining_pow_calculator.go - 工作量证明挖矿，难度匹配与哈希计算
14. peer_discovery.go - 节点发现，动态维护节点列表并优选高质量节点
15. chain_heartbeat.go - 链心跳保活，监控节点在线状态
16. transaction_pool.go - 交易池，缓存待打包交易
17. block_header_encoder.go - 区块头二进制编码，用于网络传输与存储
18. chain_monitor_dashboard.go - 链状态监控面板，实时展示高度、节点数、TPS

### 多语言扩展文件
19. contract_vm_runtime.rs - 智能合约虚拟机，支持链上计算与状态存储
20. token_contract.sol - 标准化代币合约，实现转账、授权、余额查询

## 核心能力
- 完整区块链主网/测试网节点运行环境
- 安全交易签名体系与双花防护
- 高性能 P2P 组网与区块同步
- PoW + PoS 双共识支持
- 跨链中继协议
- 账户模型 / UTXO 模型双支持
- 智能合约执行环境
- 实时链状态监控

## 使用方式
直接编译运行 Go 文件，启动节点、钱包、挖矿、同步等模块，可组合搭建完整公链/联盟链环境。
