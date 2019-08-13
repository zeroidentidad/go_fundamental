let ws = null

const setSystemMessage = data => {
    systemMessage.textContent = data
}

const login = async() => {
    const user = {
        name: usrName.value,
        password: password.value
    }

    const header = new Headers()
    header.append('Content-Type', 'application/json')
    let data = {}
    const response = await fetch('/login', { method: 'POST', headers: header, body: JSON.stringify(user) })
    switch (response.status) {
        case 200:
           data = await response.json()
            connectWS(data)
            console.log(data)
            setSystemMessage(`Conectado OK. Code: ${response.status}`)
            break;
        case 401:
            setSystemMessage(`Usuario o contraseña no valido. Code: ${response.status}`)
            break;            
        default:
            setSystemMessage(`Estado no esperado. Code: ${response.status}`);
            break;
    }
}

btnLogin.addEventListener('click', e => {
    e.preventDefault()
    login()
})

const connectWS = data => {
    ws = new WebSocket(`ws://localhost:9000/ws?uname=${usrName.value}&token=${data.token}`)
    ws.onopen = e => {
        setSystemMessage('Conectado a websocket correctamente.')
    }
    ws.onerror = e => {
        setSystemMessage(`Ocurrio un error. Desc: ${e}`)
    }
    ws.onmessage = e => {
        const data = JSON.parse(e.data)
        switch (data.type) {
            case 'message':
                content.insertAdjacentHTML('beforeend',`<div>De: ${data.data_response.name}, Mensaje: ${data.data_response.message}</div>`)
                break;
            case 'sale':

                break;
            case 'pong':

                break;                        
            default:
                setSystemMessage('Mensaje desconocido recibido')
                break;
        }
    }
}

btnSend.addEventListener('click', e => {
    e.preventDefault()
    let data={};
    if (e.data === null || e.data ===''){
        setSystemMessage('Mensaje vacio.')        
    }else{
        data = {
            type: 'message',
            message: textMsg.value
        }
        ws.send(JSON.stringify(data))
    }
})