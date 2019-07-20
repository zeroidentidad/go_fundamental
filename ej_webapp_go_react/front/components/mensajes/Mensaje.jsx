import React, { Component } from 'react';
import PropTypes from 'prop-types';
import fecha from 'fecha';

class Mensaje extends Component {
    render() {
        let {mensaje} = this.props;
        let createdAt = fecha.format(mensaje.createdAt, 'HH:mm:ss MM/DD/YYYY')
        return (
            <li className='message'>
                <div className='author'>
                    <strong>{mensaje.autor}</strong>
                    <i className='timestamp'>{createdAt}</i>
                </div>
                <div className='body'>{mensaje.body}</div>
            </li>
        )
    }
}

Mensaje.propTypes = {
    mensaje: PropTypes.object.isRequired
}

export default Mensaje;