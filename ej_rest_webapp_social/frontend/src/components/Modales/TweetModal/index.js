import React, {useState} from "react";
import { Modal, Form, Button } from "react-bootstrap";
import classNames from "classnames";
import { toast } from "react-toastify";
import { Close } from "../../../utils/icons";
import { addTweetApi } from "../../../api/tweet";

import "./TweetModal.scss";

export default function TweetModal(props) {
    const { show, setShow } = props;
    const [message, setMessage] = useState("");
    const maxLength = 280;

    const onSubmit = (e) => {
        e.preventDefault();
        if (message.length > 0 && message.length <= maxLength) {
        addTweetApi(message)
            .then((response) => {
            if (response?.code >= 200 && response?.code < 300) {
                toast.success(response.message);
                setShow(false);
                window.location.reload();
            }
            })
            .catch(() => {
            toast.warning("Error al enviar tuit");
            });
        }
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