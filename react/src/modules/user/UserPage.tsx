import React, {useEffect, useState} from "react";
import {useLoaderData, useNavigate} from "react-router-dom";
import {FetchUserQuery, User, UserBilling} from "../../__generated__/graphql";
import Panel from "../../common/components-library/Panel";
import Spinner from "../../common/components/Spinner";
import {Button, Modal, useForm, useModal} from "../../common/components/shelly-ui";
import {FontAwesomeIcon} from "@fortawesome/react-fontawesome";
import {faPencil} from "@fortawesome/free-solid-svg-icons";
import UserBillingForm from "./components/UserBillingForm";

const UserPage: React.FC = () => {
    const res = useLoaderData() as { data: FetchUserQuery, loading: boolean };
    const [user, setUser] = useState<User | null>(null);
    const [isLoading, setLoading] = useState<boolean>(false);
    const navigate = useNavigate();

    useEffect(() => {
        if (res?.data) {
            setUser(res.data.user as User);
        }

        setLoading(res.loading);
    }, [res]);

    const billingModal = useModal();
    const billingForm = useForm({
        onSuccess: () => {
            billingModal.close();
            navigate( '.', {replace: true} );
        }
    });

    if (isLoading || !user) {
        return <main className="relative">
            <Spinner/>
        </main>;
    }

    return <main className="flex flex-wrap gap-2 w-full">
        <Modal modal={billingModal}>
            <Modal.Title>
                Modifica i dati di fatturazione
            </Modal.Title>
            <UserBillingForm form={billingForm} userBilling={user.userBilling as UserBilling}/>
            <Modal.Actions>
                <Button onClick={billingModal.close}>
                    Annulla
                </Button>
                <Button buttonType="primary" onClick={billingForm.submitForm}>
                    Salva
                </Button>
            </Modal.Actions>
        </Modal>
        <Panel className="flex-auto w-1/2">
            <Panel.Title>
                Dati dell'utente
            </Panel.Title>
            <ul>
                <li><strong>Nome: </strong> {user.name}</li>
                <li><strong>Cognome: </strong> {user.surname}</li>
                <li><strong>Email: </strong> {user.email}</li>
            </ul>
        </Panel>
        {
            user.userBilling &&
            <>
                <Panel className="flex-auto w-1/2">
                    <Panel.Title>
                        Dati di fatturazione
                        <Button size="sm" onClick={() => billingModal.open()}>
                            <FontAwesomeIcon icon={faPencil} />
                        </Button>
                    </Panel.Title>
                    <ul>
                        <li><strong>Nome:</strong> {user.userBilling.name}</li>
                        <li><strong>Cognome:</strong> {user.userBilling.surname}</li>
                        <li><strong>Indirizzo 1:</strong> {user.userBilling.addressLine1}</li>
                        {user.userBilling.addressLine2 &&
                            <li><strong>Indirizzo 2:</strong> {user.userBilling.addressLine2}</li>}
                        <li><strong>Città:</strong> {user.userBilling.city}</li>
                        <li><strong>Paese:</strong> {user.userBilling.country}</li>
                        <li><strong>Provincia:</strong> {user.userBilling.province}</li>
                        <li><strong>CAP:</strong> {user.userBilling.cap}</li>
                        <li><strong>Codice fiscale:</strong> {user.userBilling.fiscalCode}</li>
                        <li><strong>Partita IVA:</strong> {user.userBilling.vatNumber}</li>
                        <li><strong>Tipo entità legale:</strong> {user.userBilling.legalEntityType.name}</li>
                    </ul>
                </Panel>
                <Panel className="flex-auto">
                    <Panel.Title>
                        Fatturazione elettronica
                    </Panel.Title>
                    <ul>
                        <li><strong>Codice destinatario: </strong> {user.userBilling?.sdiCode}</li>
                        <li><strong>PEC: </strong> {user.userBilling.sdiPec}</li>
                    </ul>
                </Panel>
            </>
        }
    </main>;
};

export default UserPage;