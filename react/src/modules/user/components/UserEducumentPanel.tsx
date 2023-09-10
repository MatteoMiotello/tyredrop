import React from "react";
import {UserBilling} from "../../../__generated__/graphql";
import Panel from "../../../common/components-library/Panel";
import {Button, Modal, useForm, useModal} from "../../../common/components/shelly-ui";
import UserEdocumentForm from "./UserEdocumentForm";
import {FontAwesomeIcon} from "@fortawesome/react-fontawesome";
import {faPencil} from "@fortawesome/free-solid-svg-icons";


type UserEdocumentPanelProps = {
    userBilling: UserBilling
    className?: string
    onUpdate?: () => void
}
const UserEducumentPanel: React.FC<UserEdocumentPanelProps> = ( {userBilling, className, onUpdate}) => {
    const modal = useModal();
    const form = useForm({
        onSuccess: () => {
            modal.close();
            if (onUpdate) {
                onUpdate();
            }
        }
    });

    return <>
        <Modal modal={modal}>
            <Modal.Title>
                Modifica i dati di fatturazione elettronica
            </Modal.Title>
            <UserEdocumentForm form={form} userBilling={userBilling}/>
            <Modal.Actions>
                <Button onClick={modal.close}>
                    Annulla
                </Button>
                <Button buttonType="primary" onClick={ form.submitForm }>
                    Salva
                </Button>
            </Modal.Actions>
        </Modal>
        <Panel className={className}>
        <Panel.Title>
            Fatturazione elettronica
            <Button size="sm" onClick={modal.open}>
                <FontAwesomeIcon icon={faPencil}/>
            </Button>
        </Panel.Title>
        <ul>
            <li><strong>Codice destinatario: </strong> {userBilling.sdiCode}</li>
            <li><strong>PEC: </strong> {userBilling.sdiPec}</li>
        </ul>
    </Panel>
        </>;
};

export default UserEducumentPanel;