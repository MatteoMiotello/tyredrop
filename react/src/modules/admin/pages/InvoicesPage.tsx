import {faPlus} from "@fortawesome/free-solid-svg-icons";
import {FontAwesomeIcon} from "@fortawesome/react-fontawesome";
import React, {useEffect, useState} from "react";
import {
    AllInvoicesQuery,
    AllInvoicesQueryVariables,
    DeleteInvoiceMutation, DeleteInvoiceMutationVariables,
    Invoice
} from "../../../__generated__/graphql";
import {useMutation, useQuery} from "../../../common/backend/graph/hooks";
import {DELETE_INVOICE} from "../../../common/backend/graph/mutation/invoice";
import {ALL_INVOICES} from "../../../common/backend/graph/query/invoice";
import Panel from "../../../common/components-library/Panel";
import {Button, Modal, useForm, useModal} from "../../../common/components/shelly-ui";
import ConfirmModal from "../../user/components/ConfirmModal";
import InvoiceAdminForm from "../components/Invoice/InvoiceAdminForm";
import InvoiceAdminTable from "../components/Invoice/InvoiceAdminTable";

const InvoicesPage: React.FC = () => {
    const query = useQuery<AllInvoicesQuery, AllInvoicesQueryVariables>(ALL_INVOICES, {
        variables: {
            pagination: {
                limit: 20,
                offset: 0
            },
            input: {}
        },
        fetchPolicy: 'no-cache'
    });
    const [mutate] = useMutation<DeleteInvoiceMutation, DeleteInvoiceMutationVariables>( DELETE_INVOICE );

    const [invoiceToDelete, setInvoiceToDelete] = useState<Invoice | undefined>();

    const modal = useModal();
    const form = useForm({
        onSuccess: () => {
            modal.close();
            query.refetch();
        }
    });

    const confirmModal  = useModal({
        onClose: () => setInvoiceToDelete(undefined)
    });

    useEffect(() => {
        if ( invoiceToDelete === undefined ) {
            return;
        }

        confirmModal.open();
    }, [invoiceToDelete]);

    return <main>
        <ConfirmModal modal={confirmModal}
                      onConfirm={ () => {
                          if ( invoiceToDelete === undefined ) {
                              return;
                          }

                          mutate( {
                              variables: {
                                  id: invoiceToDelete.id
                              }
                          } ).then( () => {
                              confirmModal.close();
                              query.refetch();
                          } );
                      } }
                      modalTitle="Confermi di voler eliminare la fattura?"/>
        <Modal modal={modal}>
            <Modal.Title>
                Inserisci una nuova fattura
            </Modal.Title>
            <InvoiceAdminForm form={form}/>
            <Modal.Actions>
                <Button onClick={modal.close}>
                    Annulla
                </Button>
                <Button buttonType="primary" onClick={form.submitForm}>
                    Salva
                </Button>
            </Modal.Actions>
        </Modal>
        <Panel>
            <Panel.Title>
                Tutte le fatture
                <Button onClick={modal.open}>
                    <FontAwesomeIcon icon={faPlus}/>
                </Button>
            </Panel.Title>
            {
                query.data &&
                <InvoiceAdminTable setInvoiceToDelete={setInvoiceToDelete} query={query}/>
            }
        </Panel>
    </main>;
};

export default InvoicesPage;