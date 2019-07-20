import React, { Component } from 'react';
import PropTypes from 'prop-types';
import Mensaje from './Mensaje.jsx';

class ListaMensajes extends Component {
    render() {
        return (
            <ul>
                {
                    this.props.mensajes.map(mensaje => {
                        return (
                            <Mensaje key={mensaje.id} mensaje={mensaje} />
                        )
                    })
                }
            </ul>
        )
    }
}

ListaMensajes.propTypes = {
    mensajes: PropTypes.array.isRequired
}

export default ListaMensajes;