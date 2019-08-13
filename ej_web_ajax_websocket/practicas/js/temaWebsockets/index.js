let ws = null

const setText = data => {
    const msg = `<div>${data}</div>`
    log.insertAdjacentHTML('beforeend', msg)
}

const setMensaje = data => {
    const msg = `<div><span>${data.nombre}: </span><span> ${data.mensaje}</span></div>`
    log.insertAdjacentHTML('beforeend', msg)
}

btnConn.addEventListener('click', e => {
    ws = new WebSocket('ws://demos.kaazing.com/echo');
    ws.onopen = () => setText('Conectado')
    ws.onclose = () => setText('Desconectado')
    ws.onerror = e => setText(e)
    ws.onmessage = e => {
        const mensaje = JSON.parse(e.data)
        setMensaje(mensaje)
    }
})

btnDesc.addEventListener('click', e => {
    ws.close()
})

btnEnv.addEventListener('click', e => {
    const mensaje = {
        nombre: txtNombre.value,
        mensaje: txtMsg.value
    }
    ws.send(JSON.stringify(mensaje))
})