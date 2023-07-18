import React, { Fragment } from "react";

const HomePage: React.FC = () => {

    return (
        <Fragment>
            <h2 className="container text-xs tracking-wide text-gray-500 dark:text-gray-300 uppercase font-bold mb-8">
                Monitors
            </h2>
            <div className="monitors space-y-6">
                <div className="monitor py-8 bg-green-400 bg-opacity-10" x-data="bars()">
                    <div className="container flex items-center justify-between mb-3">
                        <h3 className="text-2xl text-gray-800 dark:text-gray-100">Mastodon Web Interface</h3>
                        <span className="text-green-600 dark:text-green-400 font-semibold">
                            Operational
                        </span>
                    </div>
                    <div className="container bars">
                        <div className="flex space-x-px">
                            <template>
                                <div className="bars bg-green-400 flex-1 h-10 rounded hover:opacity-80 cursor-pointer">
                                </div>
                            </template>
                        </div>
                    </div>
                    <div className="container mt-2">
                        <div className="flex items-center">
                            <span className="pr-2 flex-shrink-0 text-green-500 text-xs font-semibold"
                                x-text="count + ' days ago'"></span>
                            <div className="h-px bg-green-500 w-full"></div>
                            <span className="px-2 flex-shrink-0 text-green-500 text-xs font-semibold">100%</span>
                            <div className="h-px bg-green-500 w-full"></div>
                            <span className="pl-2 flex-shrink-0 text-green-500 text-xs font-semibold">Today</span>
                        </div>
                    </div>
                </div>
            </div>
        </Fragment>
    );
}

export default HomePage;