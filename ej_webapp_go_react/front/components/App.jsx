import React, { Component } from 'react';
import SeccionCanales from './canales/SeccionCanales.jsx'

class App extends Component {
    constructor(props){
        super(props);
        this.state = {
            canales: []
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
        // obtener mensajes de canales
    }
    render() {
        return (
            <SeccionCanales
            canales={this.state.canales}
            agregarCanal={this.agregarCanal.bind(this)}
            setCanal={this.setCanal.bind(this)}
            />
        )
    }
}

export default App;