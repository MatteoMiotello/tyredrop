import React from "react";
import {useAuth} from "../auth/hooks/useAuth";
import {useQuery} from "@apollo/client";
import {USER_BILLING} from "../../common/backend/graph/query/users";
import Spinner from "../../common/components/Spinner";

const UserPage: React.FC = () => {
    const auth = useAuth();
    const {data, loading, error} = useQuery(USER_BILLING);

    if (loading) {
        return <Spinner></Spinner>;
    }

    console.log(data);
    return <main className="p-4">
        <h3 className="text-xl font-semibold divider"> Dati dell'utente </h3>
        <ul>
            <li><strong>Nome: </strong> {auth.user?.user?.name}</li>
            <li><strong>Cognome: </strong> {auth.user?.user?.surname}</li>
            <li><strong>Email: </strong> {auth.user?.user?.email}</li>
        </ul>
        <h3 className="text-xl font-semibold divider">Dati di fatturazione</h3>
        <ul>
            <li><strong>Nome:</strong> {data.userBilling.name}</li>
            <li><strong>Cognome:</strong> {data.userBilling.surname}</li>
            <li><strong>Indirizzo 1:</strong> {data.userBilling.addressLine1}</li>
            {data.userBilling.addressLine2 && <li><strong>Indirizzo 2:</strong> {data.userBilling.addressLine2}</li>}
            <li><strong>Città:</strong> {data.userBilling.city}</li>
            <li><strong>Paese:</strong> {data.userBilling.country}</li>
            <li><strong>Provincia:</strong> {data.userBilling.province}</li>
            <li><strong>CAP:</strong> {data.userBilling.cap}</li>
            <li><strong>Codice fiscale:</strong> {data.userBilling.fiscalCode}</li>
            <li><strong>Partita IVA:</strong> {data.userBilling.vatNumber}</li>
            <li><strong>Tipo entità legale:</strong> {data.userBilling.legalEntityType.name}</li>
        </ul>
    </main>;
};

export default UserPage;