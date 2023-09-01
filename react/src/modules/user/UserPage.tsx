import React, {useEffect, useState} from "react";
import {useLoaderData} from "react-router-dom";
import {FetchUserQuery, User} from "../../__generated__/graphql";
import Panel from "../../common/components-library/Panel";
import Spinner from "../../common/components/Spinner";

const UserPage: React.FC = () => {
    const res = useLoaderData() as { data: FetchUserQuery, loading: boolean };
    const [user, setUser] = useState<User | null>(null);
    const [isLoading, setLoading] = useState<boolean>(false);

    useEffect(() => {
        if (res?.data) {
            setUser(res.data.user as User);
        }

        setLoading(res.loading);
    }, [res]);

    if (isLoading || !user) {
        return <main className="relative">
            <Spinner/>
        </main>;
    }

    return <main className="flex flex-wrap gap-2 w-full">
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