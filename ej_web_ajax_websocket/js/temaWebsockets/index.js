let ws = null

const setText = data => {
    const msg = `<div>${data}</div>`
    log.insertAdjacentHTML('beforeend', msg)
}

btnConn.addEventListener('click', e => {
    ws = new WebSocket('ws://demos.kaazing.com/echo');
    ws.onopen = () => setText('Conectado')
    ws.onclose = () => setText('Desconectado')
    ws.onerror = e => setText(e)
    ws.onmessage = e => {
        setText(e.data)
    }
})

btnDesc.addEventListener('click', e => {
    ws.close()
})

btnEnv.addEventListener('click', e => {
    ws.send(txtMsg.value)
})