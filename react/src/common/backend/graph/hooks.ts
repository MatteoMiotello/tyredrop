import {
    ApolloCache,
    DefaultContext,
    DocumentNode,
    MutationHookOptions,
    MutationTuple,
    OperationVariables,
    QueryHookOptions,
    QueryResult,
    TypedDocumentNode
} from "@apollo/client";
import {useToast} from "../../../store/toast";
import {useTranslation} from "react-i18next";
import {useMutation as useM, useQuery as useQ} from "@apollo/client";

type QueryOptions = {
    successMessage?: string
}

export const useQuery = <TData = any, TVariables extends OperationVariables = OperationVariables>(query: DocumentNode | TypedDocumentNode<TData, TVariables>, options?: QueryOptions & QueryHookOptions<TData, TVariables>): QueryResult<TData, TVariables> => {
    const {error} = useToast();
    const {t} = useTranslation();

    const newOptions = {
        ...options,
        onError: (err) => {
            if (options?.onError) {
                options.onError(err);
            }

            let message = options?.successMessage;

            if (!message) {
                message = t('common.mutation_error_message') as string;
            }

            error(message);
        }
    } as QueryHookOptions<TData, TVariables>;

    return useQ(query, newOptions);
};

type MutationOptions = {
    successMessage?: string
    errorMessage?: string
};

export const useMutation = <TData = any, TVariables = OperationVariables, TContext = DefaultContext, TCache extends ApolloCache<any> = ApolloCache<any>>(mutate: DocumentNode | TypedDocumentNode<TData, TVariables>, options?: MutationOptions & MutationHookOptions<TData, TVariables, TContext, TCache>): MutationTuple<TData, TVariables, TContext, TCache> => {
    const {success, error} = useToast();
    const {t} = useTranslation();

    const newOptions = {
        ...options,
        onCompleted: (data, opt) => {
            if (options?.onCompleted) {
                options.onCompleted(data, opt);
            }

            let message = options?.successMessage;

            if (!message) {
                message = t('common.mutation_success_message') as string;
            }

            success(message);
        },
        onError: (data, opt) => {
            if (options?.onError) {
                options.onError(data, opt);
            }

            let message = options?.errorMessage;

            if (!message) {
                message = t('common.mutation_error_message') as string;
            }

            error(message);
        }
    } as MutationHookOptions<TData, TVariables, TContext, TCache>;


    return useM(mutate, newOptions);
};