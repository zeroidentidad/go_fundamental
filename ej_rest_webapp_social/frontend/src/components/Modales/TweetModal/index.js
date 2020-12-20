import React, {useState} from "react";
import { Modal, Form, Button } from "react-bootstrap";
import { Close } from "../../../utils/icons";

import "./TweetModal.scss";

export default function TweetModal(props) {
    const { show, setShow } = props;
    const [message, setMessage] = useState("");
    const maxLength = 280;

    const onSubmit = (e) => {
        e.preventDefault();
        console.log("enviando");
    }
    
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
                <Form onSubmit={onSubmit}>
                <Form.Control
                    as="textarea"
                    rows="6"
                    placeholder="Qué pedo?"
                    onChange={(e) => setMessage(e.target.value)}
                />
                <span
                    className={classNames("count", {
                    error: message.length > maxLength,
                    })}
                >
                    {message.length}
                </span>
                <Button
                    type="submit"
                    disabled={message.length > maxLength || message.length < 1}
                >
                    Enviar
                </Button>
                </Form>
             </Modal.Body>
        </Modal>
    )
}