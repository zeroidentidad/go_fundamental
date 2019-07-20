import React, {Component} from 'react';
import PropTypes from 'prop-types';

class Canal extends Component{
    onClick(e){
        e.preventDefault();
        const {setCanal, canal} = this.props;
        setCanal(canal);
    }
    render(){
        const {canal, canalActivo}=this.props;
        const activo = canal === canalActivo ? 'active' : '';
        return(
            <li className={activo}>
                <a onClick={this.onClick.bind(this)}>{canal.nombre}</a>
            </li>
        )
    }
}

Canal.propTypes = {
    canal: PropTypes.object.isRequired,
    setCanal: PropTypes.func.isRequired,
    canalActivo: PropTypes.object.isRequired
}

export default Canal;