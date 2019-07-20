import React, { Component } from 'react';
import PropTypes from 'prop-types';

class FormMensaje extends Component {
    onSubmit(e) {
        e.preventDefault();
        const nodo = this.refs.mensaje;
        const mensaje = nodo.value;
        this.props.agregarMensaje(mensaje);
        nodo.value = '';
    }
    render() {
        let input;
        if(this.props.canalActivo.id!==undefined){
            input = (
                <input 
                ref='mensaje'
                type='text'
                className='form-control'
                placeholder='Agregar mensaje...'
                />
            )
        }
        return (
            <form onSubmit={this.onSubmit.bind(this)}>
                <div className='form-group'>{input}</div>
            </form>
        )
    }
}

FormMensaje.propTypes = {
    canalActivo: PropTypes.object.isRequired,
    agregarMensaje: PropTypes.func.isRequired
}

export default FormMensaje;