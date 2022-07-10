const chatForm = document.getElementById('chat-form')
const chatMessage = document.querySelector('.chat-messages')
const roomName = document.getElementById('room-name')
const userList = document.getElementById('users')

const {username, room} = Qs.parse(location.search, {
    ignoreQueryPrefix: true
})

const newDate = new Date()

console.log(username, room)

ws = new WebSocket("ws://localhost:8877/ws")

ws.onopen = () => {
    console.log("ws connect success")
    let loginData = {
        type: "join",
        data: {
            username,
            room
        }
    }
    ws.send(JSON.stringify(loginData))
}

ws.onmessage = (event) => {
    console.log(event.data)
    let msg = JSON.parse(event.data)
    switch (msg.type) {
        case "join":
            outputMessage(msg.data)
            break
        case "exit":
            outputMessage(msg.data)
            break
        case "msg":
            outputMessage(msg.data)
            break
        case "room":
            outputRoomName(msg.room || '')
            outputUser(msg.users || [])
    }
   
    
}

ws.onclose = (event) => {
    console.log("connection close: ", event)
}

chatForm.addEventListener('submit', (e) => {
    e.preventDefault();

    const msg = e.target.elements.msg.value
    const msgData = {
        type: "msg",
        data: msg
    }
    ws.send(JSON.stringify(msgData))

    
    outputMessageSelf({
        username: "you",
        time: newDate.toDateString(),
        data: msg
    })

    e.target.elements.msg.value = '';
    e.target.elements.msg.focus(); 
})

function outputMessage(msg) {
    const div = document.createElement('div');
    div.classList.add('message')
    div.innerHTML = `<p class="meta"> ${msg.username} <span> ${msg.time}</span></p>
    <p class="text">
    ${msg.data}
    </p>`
    document.querySelector('.chat-messages').appendChild(div)
}

function outputMessageSelf(msg) {
    const div = document.createElement('div');
    div.classList.add('message')
    div.innerHTML = `<p class="meta-self"> ${msg.username} <span> </span></p>
    <p class="text">
    ${msg.data}
    </p>`
    document.querySelector('.chat-messages').appendChild(div)
}

function outputRoomName(room) {
    roomName.innerText = room
}

function outputUser(users) {
    userList.innerHTML = `
    ${users.map(user => `<li>${user}</li>`).join('')}
    `
}