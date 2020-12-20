import React, {useState} from "react";
import { Modal, Form, Button } from "react-bootstrap";
import { Close } from "../../../utils/icons";

import "./TweetModal.scss";

export default function TweetModal(props) {
    const { show, setShow } = props;
    
    return (
        <Modal className="tweet-modal"
        show={show}
        onHide={() => setShow(false)}
        centered
        size="lg"
        >
            <Modal.Header>
                <Modal.Title>
                    <Close onClick={() => setShow(false)} />
                </Modal.Title>
            </Modal.Header>
             <Modal.Body>
                 <h2>Contenido</h2>
             </Modal.Body>
        </Modal>
    )
}