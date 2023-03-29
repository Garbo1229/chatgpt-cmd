### 一个简单golang命令行接入chatgpt的脚本

使用步骤：

1. 复制代码
    `git clone https://github.com/Garbo1229/chatgpt-cmd.git`

2. 修改 `main.go` 中`apiKey` 的变量的值
    ```
    ...
    // 替换为您的API密钥
	var apiKey = ""
    ...
    ```

3. 修改 `~/.bashrc` 或 `~/.zshrc`
    ```
    // 最后一行添加 PATH 为你的路径
    ...
    alias chatgpt="go run <PATH>/main.go"
    ```

4. 立即生效
    ```
    source ~/.bashrc 或 ~/.zshrc
    ```

5. 测试
    `chatgpt`

Enjoy it ~