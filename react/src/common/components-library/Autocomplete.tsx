import {faCheck, faChevronDown} from "@fortawesome/free-solid-svg-icons";
import {FontAwesomeIcon} from "@fortawesome/react-fontawesome";
import {Combobox, Transition} from "@headlessui/react";
import React, {Fragment, useEffect, useState} from "react";

export type AutocompleteQueryHandler = (query: string) => Promise<AutocompleteOption[]>

type AutocompleteProps = {
    getOptions: AutocompleteQueryHandler
    initialOptions: AutocompleteOption[]
    className?: string
    name: string
    placeholder?: string | undefined
}

type AutocompleteOption = {
    title: string,
    value: any
}

const Autocomplete: React.FC<AutocompleteProps> = (props) => {
    const [selected, setSelected] = useState(props.initialOptions[0] ?? null);
    const [query, setQuery] = useState('');
    const [filteredOptions, setFilteredOptions] = useState(props.initialOptions);

    useEffect(() => {
        props.getOptions(query).then(res => {
            setFilteredOptions(res);
        })
            .catch((err) => {
                return null;
        });
    }, [query]);

    return <Combobox value={selected} onChange={setSelected} name={props.name}>
        <div className={"relative " + props.className}>
            <label
                className="select relative w-full cursor-default overflow-hidden p-0">
                <Combobox.Input<AutocompleteOption>
                    className="input w-full border-none p-4"
                    displayValue={(option) => option ? option.title : ''}
                    onChange={(event) => setQuery(event.target.value)}
                    placeholder={props.placeholder}
                />
                <Combobox.Button className="absolute inset-y-0 right-0 flex items-center pr-2 border-none">
                    <FontAwesomeIcon icon={faChevronDown}
                                     className="h-5 w-5 text-primary label-text"
                                     aria-hidden="true"
                    />
                </Combobox.Button>
            </label>
            <Transition
                as={Fragment}
                leave="transition ease-in duration-100"
                leaveFrom="opacity-100"
                leaveTo="opacity-0"
                afterLeave={() => setQuery('')}
            >
                <Combobox.Options
                    className="absolute mt-1 pl-0 w-full overflow-auto rounded-md bg-white py-1 text-base shadow-lg ring-1 ring-black ring-opacity-5 focus:outline-none sm:text-sm">
                    {filteredOptions.length === 0 && query !== '' ? (
                        <div className="relative cursor-default select-none py-2 px-4 text-gray-700">
                            Nothing found.
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