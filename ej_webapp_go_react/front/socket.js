import {EventEmitter} from 'events';

class Socket {
    constructor(ws=new WebSocket(), ee = new EventEmitter()){
        //super();
        this.ws = ws;
        this.ee = ee;
        ws.onmessage = this.message.bind(this);
        ws.onopen = this.open.bind(this);
        ws.onclose = this.close.bind(this);        
    }
    on(nombre, fn){
        this.ee.on(nombre, fn);
    }
    off(nombre, fn) {
        this.ee.removeListener(nombre, fn);
    }
    emit(nombre, data){
        const mensaje = JSON.stringify({ nombre, data});
        this.ws.send(mensaje);
    }    
    // handlers websocket test
    message(e) {
        try {
            const mensaje = JSON.parse(e.data);
            this.ee.emit(mensaje.nombre, mensaje.data);           
        } catch (err) {
            this.ee.emit('error', err);
        }
    }
    open() {
        //this.setState({ connected: true });
        this.ee.emit('connect');
    }
    close() {
        //this.setState({ connected: false });
        this.ee.emit('disconnect');
    }    
}

export default Socket;