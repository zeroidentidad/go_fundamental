import React, { Component } from 'react';
import PropTypes from 'prop-types';

class FormUsuario extends Component {
    onSubmit(e){
        e.preventDefault();
        const nodo = this.refs.nombreUsuario;
        const nombreUsuario = nodo.value;
        this.props.setNombreUsuario(nombreUsuario);
        nodo.value='';
    }
    render() {
        return (
            <form onSubmit={this.onSubmit.bind(this)}>
                <div className='form-group'>
                    <input
                    ref='nombreUsuario'
                    type='text'
                    className='form-control'
                    placeholder='Ingresa tu nombre...'
                    />
                </div>
            </form>
        )
    }
}

FormUsuario.propTypes = {
    setNombreUsuario: PropTypes.func.isRequired
}

export default FormUsuario;