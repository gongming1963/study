document.addEventListener('DOMContentLoaded', function() {
    // 页面加载完成后自动请求欢迎语
    requestWelcomeMessage();
});

document.getElementById('send-btn').addEventListener('click', sendMessage);

document.getElementById('chat-input').addEventListener('keypress', function(event) {
    if (event.keyCode === 13) {
        event.preventDefault();
        sendMessage();
    }
});

function sendMessage() {
    var inputElement = document.getElementById('chat-input');
    var inputValue = inputElement.value.trim();
    if (inputValue !== '') {
        addMessageToChat('User', inputValue);
        callChatbotApi(inputValue);
        inputElement.value = '';
    }
}

function addMessageToChat(sender, message) {
    var chatOutput = document.getElementById('chat-output');
    chatOutput.innerHTML += '<div>' + sender + ': ' + message + '</div>';
}

function callChatbotApi(message) {
    fetch('http://127.0.0.1:3000/chat/', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({ message: message })
    })
    .then(response => response.json())
    .then(data => {
        // 假设API返回的数据中包含字段'reply'作为机器人的回复
        addMessageToChat('小依', data.reply);
    })
    .catch(error => {
        console.error('Error calling the chatbot API:', error);
        addMessageToChat('小依', 'Sorry, I am having trouble responding right now.');
    });
}

function requestWelcomeMessage() {
    // 这里假设发送一个空消息或特定消息会触发欢迎语
    callChatbotApi('[reset]'); // 或者您可以发送一个特定的欢迎消息请求
}
