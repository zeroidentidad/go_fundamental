import React, {Component} from 'react';
import PropTypes from 'prop-types';

class Canal extends Component{
    onClick(e){
        e.preventDefault();
        const {setCanal, canal} = this.props;
        setCanal(canal);
    }
    render(){
        const {canal}=this.props;
        return(
            <li>
                <a onClick={this.onClick.bind(this)}>{canal.nombre}</a>
            </li>
        )
    }
}

Canal.propTypes = {
    canal: PropTypes.object.isRequired,
    setCanal: PropTypes.func.isRequired
}

export default Canal;