btn.addEventListener('click', () => {loadComics()});

const loadComics = async()=>{
    const hash = 'f868e45dbbd38a5d6de3ed86a92f4d4b', apikey = '42de5550980de3333222524000d53b99';
    const nombre = ipt.value;
    const url = `https://gateway.marvel.com:443/v1/public/characters?ts=1&apikey=${apikey}&hash=${hash}&limit=20&nameStartsWith=${nombre}`;

    //Await Fetch
    const res = await fetch(url)
    switch (res.status) {
        case 200:
            const res_data = await res.json()
            dibujar(res_data.data.results);
            break;
        default:
            setText(`Error desconocido. Code: ${res.status}`);
            break;
    }
}

const dibujar = data => {
    while (miContenido.firstChild) {
        miContenido.removeChild(miContenido.firstChild);
    }

    const contenedor = document.createElement('div');
    data.forEach(comic => {
        const comicHTML = 
        `<div>
        <h2>${comic.name}</h2>
        <img src="${comic.thumbnail.path}/portrait_incredible.${comic.thumbnail.extension}" alt="${comic.name}"/>
        <p>${comic.description}</p>
        </div>`

        contenedor.insertAdjacentHTML('beforeend', comicHTML);
    });

    miContenido.appendChild(contenedor);
}

const setText = data => {
    miContenido.innerHTML = data;
}