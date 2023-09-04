import React from "react";
import Panel from "../../../common/components-library/Panel";
import {Button, Modal, useForm, useModal} from "../../../common/components/shelly-ui";
import {FontAwesomeIcon} from "@fortawesome/react-fontawesome";
import {faPencil} from "@fortawesome/free-solid-svg-icons";
import {UserBilling} from "../../../__generated__/graphql";
import UserBillingForm from "./UserBillingForm";

type UserBillingPanelProps = {
    userBilling: UserBilling
    className?: string
    onUpdate?: () => void
}
const UserBillingPanel: React.FC<UserBillingPanelProps> = ( {userBilling, className, onUpdate} ) => {
    const billingModal = useModal();
    const billingForm = useForm({
        onSuccess: () => {
            billingModal.close();
            if (onUpdate) {
                onUpdate();
            }
        }
    });

    return <>
        <Modal modal={billingModal}>
            <Modal.Title>
                Modifica i dati di fatturazione
            </Modal.Title>
            <UserBillingForm form={billingForm} userBilling={userBilling as UserBilling}/>
            <Modal.Actions>
                <Button onClick={billingModal.close}>
                    Annulla
                </Button>
                <Button buttonType="primary" onClick={billingForm.submitForm}>
                    Salva
                </Button>
            </Modal.Actions>
        </Modal>
        <Panel className={className}>
        <Panel.Title>
            Dati di fatturazione
            <Button size="sm" onClick={() => billingModal.open()}>
                <FontAwesomeIcon icon={faPencil} />
            </Button>
        </Panel.Title>
        <ul>
            <li><strong>Nome completo:</strong> {userBilling.name} {userBilling.surname}</li>
            <li><strong>Indirizzo 1:</strong> {userBilling.addressLine1}</li>
            {userBilling.addressLine2 &&
                <li><strong>Indirizzo 2:</strong> {userBilling.addressLine2}</li>}
            <li><strong>Città:</strong> {userBilling.city}</li>
            <li><strong>Paese:</strong> {userBilling.country}</li>
            <li><strong>Provincia:</strong> {userBilling.province}</li>
            <li><strong>CAP:</strong> {userBilling.cap}</li>
            <li><strong>Codice fiscale:</strong> {userBilling.fiscalCode}</li>
            <li><strong>Partita IVA:</strong> {userBilling.vatNumber}</li>
            <li><strong>Tipo entità legale:</strong> {userBilling.legalEntityType.name}</li>
        </ul>
    </Panel>
        </>;
};

export default UserBillingPanel;