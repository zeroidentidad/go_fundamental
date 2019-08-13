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
           //connectWS(data)
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