//Ejemplo simple: promesas en secuencia

const setText = data => {
    miContenido.innerHTML = data;
}

const getData = () => {
    return new Promise((resolver, rechazar) => {
        setText('<h2>Solicitando autorización...</h2><br/>');
        setTimeout(() => {
            resolver(true);
        }, 2500)
    })
}

const showData = () => {
    return new Promise((resolver, rechazar) => {
        setText('<h2>Esperando información...</h2><br/>');
        setTimeout(() => {
            resolver({nombre: 'Jesus'});
        }, 2500)
    })
}

btn.addEventListener('click', () => {
    getData().then(autorizar => {
        //return true por default
        if (autorizar) {
            return showData()
        }
    }).then(usuario => {
        setText(`<h2>${usuario.nombre}</h2><br/>`)
    })
})