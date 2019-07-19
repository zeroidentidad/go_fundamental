import React, { Component } from 'react';
import PropTypes from 'prop-types';

class FormCanal extends Component {
    onSubmit(e){
        e.preventDefault();
        const nodo = this.refs.canal;
        const nombreCanal = nodo.value;
        this.props.agregarCanal(nombreCanal);
        nodo.value='';
    }
    render() {
        return (
            <form onSubmit={this.onSubmit.bind(this)}>
                <input type='text' ref='canal'/>
            </form>
        )
    }
}

FormCanal.propTypes = {
    agregarCanal: PropTypes.func.isRequired
}

export default FormCanal;