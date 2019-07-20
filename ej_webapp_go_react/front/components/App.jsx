import React, { Component } from 'react';
import SeccionCanales from './canales/SeccionCanales.jsx'
import SeccionUsuarios from './usuarios/SeccionUsuarios.jsx';
import SeccionMensajes from './mensajes/SeccionMensajes.jsx';

class App extends Component {
    constructor(props){
        super(props);
        this.state = {
            canales: [],
            canalActivo: { canalActivo: null },
            usuarios: [],
            mensajes: []
        };
    }
    agregarCanal(nombre){
        let {canales}=this.state;
        canales.push({id: canales.length, nombre});
        this.setState({canales});
        // enviar al servidor
    }
    setCanal(canalActivo){
        this.setState({canalActivo});
        // obtener canales de mensajes
    }
    setNombreUsuario(nombre){
        let {usuarios} = this.state;
        usuarios.push({id:usuarios.length, nombre});
        this.setState({usuarios});
        // enviar al servidor
    }
    agregarMensaje(body){
        let {mensajes, usuarios} = this.state;
        let createdAt = new Date;
        let autor = usuarios.length > 0 ? usuarios[0].nombre: 'anonimo';
        mensajes.push({id:mensajes.length, body, createdAt, autor});
        this.setState({mensajes});
        // enviar al servidor
    }
    render() {
        return (
            <div className='app'>
                <div className='nav'>
                    <SeccionCanales
                        /*canales={this.state.canales}*/
                        {...this.state}
                        agregarCanal={this.agregarCanal.bind(this)}
                        setCanal={this.setCanal.bind(this)}
                    />
                    <SeccionUsuarios 
                        {...this.state}
                        setNombreUsuario={this.setNombreUsuario.bind(this)}
                    />
                </div>
                <SeccionMensajes
                    {...this.state}
                    agregarMensaje={this.agregarMensaje.bind(this)}
                />
            </div>
        )
    }
}

export default App;