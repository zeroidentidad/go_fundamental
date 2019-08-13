//Ejemplo simple: callbacks sincronas (secuenciales)

const setText = data => {
    miContenido.innerHTML = data;
}

const getData = callback =>{
    setText('<h2>Solicitando autorización...</h2><br/>');
    setTimeout(() => {
        callback(true);
    }, 2500);
}

const showData = callback => {
    setText('<h2>Esperando información...</h2><br/>');
    setTimeout(() => {
        callback({nombre: 'Jesus'});
    }, 2500);
}

btn.addEventListener('click', () => {
    getData(autorizar => {
        //return true por default
        if (autorizar){
            showData(usuario => {
                setText(`<h2>${usuario.nombre}</h2><br/>`)
            })
        }
    })
})