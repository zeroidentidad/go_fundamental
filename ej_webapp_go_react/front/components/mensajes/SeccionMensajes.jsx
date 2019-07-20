import React, { Component } from 'react';
import PropTypes from 'prop-types';
import ListaMensajes from './ListaMensajes.jsx';
import FormMensaje from './FormMensaje.jsx';

class SeccionMensajes extends Component {
    render() {
        let {canalActivo} = this.props;
        return (
            <div className='messages-container panel panel-default'>
                <div className='panel-heading'>
                    <strong>{canalActivo.nombre}</strong>
                </div>
                <div className='panel-body messages'>
                    <ListaMensajes {...this.props} />
                    <FormMensaje {...this.props} />
                </div>
            </div>
        )
    }
}

SeccionMensajes.propTypes = {
    mensajes: PropTypes.array.isRequired,
    canalActivo: PropTypes.object.isRequired,
    agregarMensaje: PropTypes.func.isRequired
}


export default SeccionMensajes;