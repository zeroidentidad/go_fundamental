
btn.addEventListener('click', e => {showMarvel()});

const showMarvel = async()=>{
    const hash = 'f868e45dbbd38a5d6de3ed86a92f4d4b';
    const apikey = '42de5550980de3333222524000d53b99';
    const nombre = ipt.value;
    const url = `https://gateway.marvel.com:443/v1/public/characters?ts=1&apikey=${apikey}&hash=${hash}&limit=20&nameStartsWith=${nombre}`;
    const req = { method: 'GET', url: url, sync:true};
    const res = await ajax(req);

    switch (res.status) {
        case 200:
            dibujar(JSON.parse(res.responseText).data.results);
            break;
        case 400:
            setText(`Error cliente web. Code: ${res.status}`);
            break;               
        default:
            setText(`Error desconocido. Code: ${res.status}`);
            break;
    }
}

const setText = data => {
    miContenido.innerHTML = data;
}

const dibujar = data => {
    while (miContenido.firstChild) {
        miContenido.removeChild(miContenido.firstChild);
    }

    const fragmento = document.createDocumentFragment();

    data.forEach(comic => {
        const contenedor = document.createElement('div');
        const titulo = document.createElement('h2');
        const imagen = document.createElement('img');
        const descripcion = document.createElement('p');

        titulo.textContent = comic.name;
        imagen.src = `${comic.thumbnail.path}/portrait_incredible.${comic.thumbnail.extension}`;
        descripcion.textContent = comic.description;

        contenedor.appendChild(titulo);
        contenedor.appendChild(imagen);
        contenedor.appendChild(descripcion);

        fragmento.appendChild(contenedor);
    });

    miContenido.appendChild(fragmento);
}