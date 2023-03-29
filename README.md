### 一个简单golang命令行接入chatgpt的脚本

使用步骤：

1. 复制代码
git clone https://github.com/Garbo1229/chatgpt-cmd.git

2. 修改apiKey的变量的值

3. 修改~/.bashrc 或 ~/.zshrc
    ```
    // 最后一行添加 PATH 为你的路径
    ...
    chatgpt (){
    go run PATH/main.go -m {$1}
    }
    ```

4. 立即生效
    ```
    source ~/.bashrc 或 ~/.zshrc
    ```

5. 测试
    ```
    chatgpt "test"
    ```

Enjoy it ~