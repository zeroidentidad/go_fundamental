//alert('Chidowan')
const cont = document.getElementById('contenido');
const btn = document.getElementById('btnAjax');
const load = document.getElementById('load');

//status
const status_ok = 200;
const status_unathorized = 401;
const status_notfound = 404;
const status_internalserver_error = 500;

load.style.display = 'none';

btn.addEventListener('click', e => {
    load.style.display = 'block';
    const xhr = new XMLHttpRequest();
    //console.log('Estado objeto: ', xhr.readyState);

    xhr.open('GET', '/json', true);
    //console.log('Estado objeto: ', xhr.readyState);

    // como usar lo solicitado
    xhr.addEventListener('load', e => {
        //cont.innerHTML = e.target.responseText;
        //console.log('Estado objeto: ', xhr.readyState);

        switch (e.target.status) {
            case status_ok:
                const data = JSON.parse(e.target.responseText);
                dibujar(data);                
                break;
            case status_unathorized:
                cont.textContent='No estas autorizado para ver el contenido';
                break;
            case status_notfound:
                cont.textContent = 'No existe, no encontrado';
                break;                
            case status_internalserver_error:
                cont.textContent = 'Error de servidor';
                break;
        }

        load.style.display = 'none';
    })

    //Peticion
    xhr.send();
    //console.log('Estado objeto: ', xhr.readyState);
})

const dibujar = data =>{

    //cont.innerHTML = '';
    while (cont.firstChild) { // el sig pasa a ser el primero
        cont.removeChild(cont.firstChild)
    }

    const fragmento = document.createDocumentFragment();

    data.forEach(nota => {
        const contenedor = document.createElement('div');
        const titulo = document.createElement('h2');
        const contenido = document.createElement('p');

        titulo.textContent = nota.Titulo;
        contenido.textContent = nota.Contenido;

        contenedor.appendChild(titulo);
        contenedor.appendChild(contenido);

        //cont.append(contenedor);
        fragmento.appendChild(contenedor);
    });

    cont.appendChild(fragmento);
}