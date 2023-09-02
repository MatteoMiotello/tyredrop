import React, {useEffect, useState} from "react";
import {useLoaderData, useNavigate} from "react-router-dom";
import Panel from "../../../common/components-library/Panel";
import {
    ChangeUserStatusMutation,
    ChangeUserStatusMutationVariables,
    FetchUserQuery, User
} from "../../../__generated__/graphql";
import {Alert, Button, useModal} from "../../../common/components/shelly-ui";
import {FontAwesomeIcon} from "@fortawesome/react-fontawesome";
import {faCheck, faTimes} from "@fortawesome/free-solid-svg-icons";
import {useMutation} from "../../../common/backend/graph/hooks";
import {UPDATE_USER_STATUS} from "../../../common/backend/graph/mutation/users";
import ConfirmModal from "../../user/components/ConfirmModal";
import UserOrdersPanel from "../components/User/UserOrdersPanel";

const UserDetailsPage: React.FC = () => {
    const data = useLoaderData() as { data: FetchUserQuery };
    const [mutate, mutationQuery] = useMutation<ChangeUserStatusMutation, ChangeUserStatusMutationVariables>(UPDATE_USER_STATUS);
    const [action, setAction] = useState<'reject' | 'confirm' | undefined>(undefined);
    const navigate = useNavigate();
    const confirmModal = useModal({
        onClose: () => setAction(undefined)
    });

    useEffect(() => {
        if (action === undefined) {
            return;
        }

        confirmModal.open();
    }, [action]);

    return <main className="flex flex-wrap gap-2 w-full p-1">
        <ConfirmModal modal={confirmModal} onConfirm={() => {
            if (action === undefined) {
                return false;
            }

            const variables: ChangeUserStatusMutationVariables = {
                userID: data.data.user?.id as string,
                confirmed: null,
                rejected: null
            };

            if (action === 'confirm') {
                variables.confirmed = true;
            } else {
                variables.rejected = true;
            }

            mutate({
                variables: variables
            }).then(() => navigate('.', {replace: true}));
        }} modalTitle={  `Confermi di voler ${ action == 'confirm' ? 'approvare' : 'rifiutare' } l'utente ${data.data.user?.email} ?`}/>
        {
            data.data?.user?.rejected &&
            <Alert type="error">
                Questo utente è stato disabilitato
            </Alert>
        }
        <Panel className="flex-auto w-1/2">
            <Panel.Title>
                <span className="w-full"> Dati dell'utente </span>
                <div className="w-full flex justify-end gap-2">
                    {
                        (!data.data.user?.rejected && !data.data.user?.confirmed) &&
                        <Button size="sm" buttonType="error" onClick={() => setAction('reject')}
                                loading={(action === 'reject' && mutationQuery.loading)}>
                            <FontAwesomeIcon icon={faTimes}/>
                            Rifiuta
                        </Button>

                    }
                    {
                        !data.data.user?.confirmed &&
                        <Button size="sm" buttonType="success" onClick={() => setAction('confirm')}
                                loading={(action === 'confirm' && mutationQuery.loading)}>
                            <FontAwesomeIcon icon={faCheck}/>
                            Conferma
                        </Button>
                    }
                </div>

            </Panel.Title>
            <ul>
                <li><strong>Nome: </strong> {data.data?.user?.name}</li>
                <li><strong>Cognome: </strong> {data.data?.user?.surname}</li>
                <li><strong>Email: </strong> {data.data?.user?.email}</li>
            </ul>
        </Panel>
        {
            data.data.user?.userBilling &&
            <>
                <Panel className="flex-auto w-1/2">
                    <Panel.Title>
                        Dati di fatturazione
                    </Panel.Title>
                    <ul>
                        <li><strong>Nome completo:</strong> {data.data.user.userBilling.name} {data.data.user.userBilling.surname}</li>
                        <li><strong>Indirizzo 1:</strong> {data.data.user.userBilling.addressLine1}</li>
                        {data.data.user.userBilling.addressLine2 &&
                            <li><strong>Indirizzo 2:</strong> {data.data.user.userBilling.addressLine2}</li>}
                        <li><strong>Città:</strong> {data.data.user.userBilling.city}</li>
                        <li><strong>Paese:</strong> {data.data.user.userBilling.country}</li>
                        <li><strong>Provincia:</strong> {data.data.user.userBilling.province}</li>
                        <li><strong>CAP:</strong> {data.data.user.userBilling.cap}</li>
                        <li><strong>Codice fiscale:</strong> {data.data.user.userBilling.fiscalCode}</li>
                        <li><strong>Partita IVA:</strong> {data.data.user.userBilling.vatNumber}</li>
                        <li><strong>Tipo entità legale:</strong> {data.data.user.userBilling.legalEntityType.name}</li>
                    </ul>
                </Panel>
                <Panel className="flex-auto">
                    <Panel.Title>
                        Fatturazione elettronica
                    </Panel.Title>
                    <ul>
                        <li><strong>Codice destinatario: </strong> {data.data.user.userBilling?.sdiCode}</li>
                        <li><strong>PEC: </strong> {data.data.user.userBilling.sdiPec}</li>
                    </ul>
                </Panel>
            </>
        }
        {
            (data.data.user?.userBilling && !data.data.user.userRole.isAdmin) &&
            <UserOrdersPanel user={data.data.user as User}/>
        }
    </main>;
};

export default UserDetailsPage;