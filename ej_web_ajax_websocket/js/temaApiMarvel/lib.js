const ajax = request => {
    return new Promise((resolve, reject)=>{
        const xhr = new XMLHttpRequest();
        xhr.open(request.method, request.url, request.sync)
        xhr.addEventListener('load', e=>{
            resolve(e.target);
        })
        xhr.send();
    })
}