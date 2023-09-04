import React, {useEffect, useState} from "react";
import {useLoaderData, useNavigate} from "react-router-dom";
import Panel from "../../../common/components-library/Panel";
import {
    ChangeUserStatusMutation,
    ChangeUserStatusMutationVariables,
    FetchUserQuery, User, UserBilling
} from "../../../__generated__/graphql";
import {Alert, Button, useModal} from "../../../common/components/shelly-ui";
import {FontAwesomeIcon} from "@fortawesome/react-fontawesome";
import {faCheck, faTimes} from "@fortawesome/free-solid-svg-icons";
import {useMutation} from "../../../common/backend/graph/hooks";
import {UPDATE_USER_STATUS} from "../../../common/backend/graph/mutation/users";
import ConfirmModal from "../../user/components/ConfirmModal";
import UserOrdersPanel from "../components/User/UserOrdersPanel";
import UserBillingPanel from "../../user/components/UserBillingPanel";
import UserEducumentPanel from "../../user/components/UserEducumentPanel";

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
        }}
                      modalTitle={`Confermi di voler ${action == 'confirm' ? 'approvare' : 'rifiutare'} l'utente ${data.data.user?.email} ?`}/>
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
                <UserBillingPanel
                    className="flex-auto w-1/2"
                    userBilling={data.data.user.userBilling as UserBilling}
                    onUpdate={() => navigate('.', {replace: true})}/>
                <UserEducumentPanel
                    className="flex-auto"
                    userBilling={data.data.user.userBilling as UserBilling}
                    onUpdate={() => navigate('.', {replace: true})}
                />
            </>
        }
        {
            (data.data.user?.userBilling && !data.data.user.userRole.isAdmin) &&
            <UserOrdersPanel user={data.data.user as User}/>
        }
    </main>;
};

export default UserDetailsPage;