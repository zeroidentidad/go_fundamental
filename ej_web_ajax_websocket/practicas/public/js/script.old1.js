//alert('Chidowan')
const cont = document.getElementById('contenido');
const btn = document.getElementById('btnAjax');
const load = document.getElementById('load');

load.style.display='none';

btn.addEventListener('click', e=>{
    load.style.display = 'block';
    const xhr = new XMLHttpRequest();
    console.log('Estado objeto: ', xhr.readyState);

    xhr.open('GET','/data.html', true);
    console.log('Estado objeto: ', xhr.readyState);

    // como usar lo solicitado
    xhr.addEventListener('load', e =>{
        cont.innerHTML=e.target.responseText;
        console.log('Estado objeto: ', xhr.readyState);
        load.style.display = 'none';
    })

    //Peticion
    xhr.send();
    console.log('Estado objeto: ', xhr.readyState);
})