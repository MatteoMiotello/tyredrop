import React, {ReactNode} from "react";
import {Tab} from "@headlessui/react";

export type Partial = {
    title: string
    content: ReactNode
}

type TabsProps = {
    parts: Partial[]
}

function classNames(...classes: string[]) {
    return classes.filter(Boolean).join(' ');
}

const Tabs: React.FC<TabsProps> = ( props) => {
    return <div className="">
        <Tab.Group>
            <Tab.List className="tabs flex justify-center">
                {props.parts.map((part, id) => (
                    <Tab
                        key={id}
                        className={({ selected }) =>
                            classNames(
                                'tab tab-bordered outline-none',
                                selected
                                    ? 'tab-active outline-none'
                                    : ''
                            )
                        }
                    >
                        {part.title}
                    </Tab>
                ))}
            </Tab.List>
            <Tab.Panels className="mt-2">
                {props.parts.map((part, idx) => (
                    <Tab.Panel key={idx} >
                        {part.content}
                    </Tab.Panel>
                ))}
            </Tab.Panels>
        </Tab.Group>
    </div>;
};

export default Tabs;