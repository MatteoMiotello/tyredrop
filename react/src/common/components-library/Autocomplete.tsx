import {faCheck, faChevronDown} from "@fortawesome/free-solid-svg-icons";
import {FontAwesomeIcon} from "@fortawesome/react-fontawesome";
import {Combobox, Transition} from "@headlessui/react";
import React, {ChangeEventHandler, Fragment, ReactNode, useEffect, useState} from "react";
import {useTranslation} from "react-i18next";
import {PropsWithValidators, ValidationHandler} from "../validation/validators";

export type AutocompleteQueryHandler = (query: string) => Promise<AutocompleteOption[] | null>

interface AutocompleteProps extends PropsWithValidators {
    getOptions: AutocompleteQueryHandler;
    initialOptions: AutocompleteOption[];
    className?: string;
    name: string;
    placeholder?: string | undefined;
    labelText?: string | undefined | null;
    defaultValue?: AutocompleteOption
}

export type AutocompleteOption = {
    title: ReactNode,
    value: any
}

const Autocomplete: React.FC<AutocompleteProps> = (props) => {
    const [selected, setSelected] = useState<AutocompleteOption | null>(props.defaultValue ?? props.initialOptions[0] ?? null);
    const [query, setQuery] = useState('');
    const [filteredOptions, setFilteredOptions] = useState(props.initialOptions);
    const [error, setError] = useState<string | null>(null);
    const {t} = useTranslation();

    useEffect( () => {
        if ( props.defaultValue ) {
            setSelected( props.defaultValue );
        }
    }, [props.defaultValue] );

    const performQuery = () => {
        const options = props.getOptions(query);

        if (!options) {
            return;
        }

        options.then(res => {
            if (res) {
                setFilteredOptions(res);
            }
        })
            .catch((err) => {
                return null;
            });
    };

    useEffect(() => {
        performQuery();
    }, [query]);

    const onChange: ChangeEventHandler<HTMLInputElement> = (event) => {
        const value = event.target.value;

        if ( !value ) {
            setSelected(null);
        }

        if (!props.validators) {
            return;
        }

        props.validators?.every((validator: ValidationHandler) => {
            const error = validator(value as string);
            setError(error);

            return !error;
        });
    };

    return <Combobox value={selected} onChange={setSelected} name={props.name}>
        <div className={"relative " + props.className}>
                {props.labelText &&
                    <label className="label">
                        <span className="label-text">
                            {props.labelText}
                        </span>
                    </label>
                }
                <div
                    className={`select select-bordered relative w-full cursor-default overflow-hidden p-0 ${error ? 'select-error' : ''}`}>
                    <Combobox.Input<AutocompleteOption>
                        autoComplete="off"
                        className={"input w-full border-none p-4 font-normal"}
                        displayValue={(option) => option ? option.title : ''}
                        onChange={(event) => {
                            setQuery(event.target.value);
                            onChange(event);
                        }}
                        onFocus={onChange}
                        placeholder={props.placeholder}
                    />
                    <Combobox.Button className="absolute inset-y-0 right-0 flex items-center pr-2 border-none" onClick={performQuery}>
                        <FontAwesomeIcon icon={faChevronDown}
                                         className="h-5 w-5 text-primary label-text"
                                         aria-hidden="true"
                        />
                    </Combobox.Button>
                </div>
            {error ? <span className="label-text-alt text-error">{error}</span> : ''}
            <Transition
                as={Fragment}
                leave="transition ease-in duration-100"
                leaveFrom="opacity-100"
                leaveTo="opacity-0"
                afterLeave={() => setQuery('')}
            >
                <Combobox.Options
                    className="absolute z-30 mt-1 pl-0 w-full overflow-auto rounded-md bg-white py-1 text-base shadow-lg ring-1 ring-black ring-opacity-5 focus:outline-none sm:text-sm">
                    {filteredOptions.length === 0 ? (
                        <div className="relative cursor-default select-none py-2 px-4 text-gray-700">
                            {query.length > 0 ? t('fields.empty_query') : t('fields.no_results')}
                        </div>
                    ) : (
                        filteredOptions.map((option, key) => (
                            <Combobox.Option
                                key={key}
                                className={({active}) =>
                                    `relative cursor-default select-none py-2 list-none ${
                                        active ? 'bg-base-200 text-secondary' : 'text-gray-900'
                                    }`
                                }
                                value={option}
                            >
                                {({selected, active}) => (
                                    <>
                        <span
                            className={`block truncate ml-8 ${
                                selected ? 'font-medium' : 'font-normal'
                            }`}
                        >
                          {option.title}
                        </span>
                                        {selected ? (
                                            <span
                                                className={`absolute inset-y-0 left-0 flex items-center text-primary ml-2 ${
                                                    active ? 'text-primary' : 'text-grey-600'
                                                }`}
                                            >
                            <FontAwesomeIcon icon={faCheck} className="h-5 w-5" aria-hidden="true"/>
                          </span>
                                        ) : null}
                                    </>
                                )}
                            </Combobox.Option>
                        ))
                    )}
                </Combobox.Options>
            </Transition>
        </div>
    </Combobox>;
};

export default Autocomplete;