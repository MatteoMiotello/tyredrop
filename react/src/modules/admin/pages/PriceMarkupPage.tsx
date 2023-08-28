import {faPlus} from "@fortawesome/free-solid-svg-icons";
import {FontAwesomeIcon} from "@fortawesome/react-fontawesome";
import React, {useEffect, useState} from "react";
import {ProductPriceMarkup} from "../../../__generated__/graphql";
import { useQuery} from "../../../common/backend/graph/hooks";
import {PRICE_MARKUPS} from "../../../common/backend/graph/query/products";
import Panel from "../../../common/components-library/Panel";
import {Button, Modal, useForm, useModal} from "../../../common/components/shelly-ui";
import PriceMarkupForm from "../components/Price/PriceMarkupForm";
import PriceMarkupTable from "../components/Price/PriceMarkupTable";

const PriceMarkupPage: React.FC = () => {
    const query = useQuery(PRICE_MARKUPS);
    const modal = useModal({
        onClose: () => {
            setMarkupToEdit(undefined);
        }
    });
    const form = useForm({
        onSuccess: () => {
            modal.close();
            query.refetch();
        }
    });

    const [markupToEdit, setMarkupToEdit] = useState<ProductPriceMarkup|undefined>();

    useEffect(() => {
        if ( !markupToEdit ) {
            return;
        }

        modal.open();
    }, [markupToEdit]);

    return <main>
        <Modal modal={modal}>
            <Modal.Title>
                Modifica un prezzo
            </Modal.Title>
            <PriceMarkupForm form={form} markup={markupToEdit}/>
            <Modal.Actions>
                <Button onClick={modal.close}>
                    Annulla
                </Button>
                <Button onClick={form.submitForm} buttonType="primary">
                    Salva
                </Button>
            </Modal.Actions>
        </Modal>
        <Panel>
            <Panel.Title>
                Tutti i prezzi
                <Button buttonType="ghost" onClick={()=>{
                    modal.open();
                }}>
                    <FontAwesomeIcon icon={faPlus}/>
                </Button>
            </Panel.Title>
            {
                query.data &&
                <PriceMarkupTable query={query} setMarkupToEdit={setMarkupToEdit}/>
            }
        </Panel>
    </main>;
};

export default PriceMarkupPage;