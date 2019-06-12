
let formCommentAdd = $('formCommentAdd'),
    commentContent = $('comment-content'),
    messageComment = $('mensaje-comentario');

formCommentAdd.addEventListener('submit', e => {
    e.preventDefault();
    let obj = {
        content: commentContent.value
    };

    peticionAJAX(formCommentAdd.method, formCommentAdd.action, JSON.stringify(obj))
    .then(respuesta =>{
        if (respuesta.status === 201) {
            messageComment.textContent = respuesta.response.message;
            //renderizar comentario
        }else{
            messageComment.textContent = respuesta.response.message;
        }
    })
    .catch(error => {
        console.log(error);
    });

});

//getComments autoinvocado anonimamente
(function(){
    peticionAJAX('GET', '/api/comments/')
    .then(respuesta =>{
        console.log(respuesta);
    })
    .catch(error => {
        console.log(error);
    });
})();