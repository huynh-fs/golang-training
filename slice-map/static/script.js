document.addEventListener('DOMContentLoaded', () => {
    const gameOutput = document.getElementById('game-output');
    const commandInput = document.getElementById('command-input');
    const sendButton = document.getElementById('send-button');

    // Hàm để thêm text vào cửa sổ game
    function appendOutput(text) {
        const p = document.createElement('p');
        p.textContent = text;
        gameOutput.appendChild(p);
        gameOutput.scrollTop = gameOutput.scrollHeight; // Cuộn xuống dưới cùng
    }

    // Hàm gửi lệnh đến server
    async function sendCommand() {
        const command = commandInput.value.trim();
        if (command === '') {
            return;
        }

        appendOutput(`> ${command}`); // Hiển thị lệnh của người chơi
        commandInput.value = ''; // Xóa input

        try {
            const response = await fetch('/command', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify({ command: command })
            });

            if (!response.ok) {
                throw new Error(`HTTP error! status: ${response.status}`);
            }

            const data = await response.json();
            appendOutput(data.output); // Hiển thị output từ server

        } catch (error) {
            console.error('Lỗi khi gửi lệnh:', error);
            appendOutput('Lỗi kết nối đến server. Vui lòng thử lại.');
        }
    }

    // Gửi lệnh khi nhấn nút
    sendButton.addEventListener('click', sendCommand);

    // Gửi lệnh khi nhấn Enter trong ô input
    commandInput.addEventListener('keypress', (e) => {
        if (e.key === 'Enter') {
            sendCommand();
        }
    });

    // Lệnh khởi tạo khi tải trang
    async function initializeGameDisplay() {
        appendOutput("Chào mừng đến với Trò Chơi Phiêu Lưu Văn Bản!");
        appendOutput("Gõ 'help' để xem các lệnh có thể sử dụng.");
        appendOutput("----------------------------------------");
        
        // Gửi lệnh 'look' ban đầu để hiển thị phòng khởi đầu
        const response = await fetch('/command', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({ command: 'look' })
        });
        const data = await response.json();
        appendOutput(data.output);
    }

    initializeGameDisplay();
});