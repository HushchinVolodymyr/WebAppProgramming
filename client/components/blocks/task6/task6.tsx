"use client";

import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import axios from "axios";
import { useState } from "react";
import { z } from "zod";

// const testData = {
//     "grinding": {
//         "n": 0.92,
//         "cos": 0.9,
//         "u": 0.38,
//         "count": 4,
//         "ph": 20,
//         "np": 80,
//         "kv": 0.15,
//         "tg": 1.33,
//         "npk": 12,
//         "npktg": 15.96,
//         "np2": 1600,
//         "efficiencyQuantity": null,
//         "activPowerCoefficient": null,
//         "estimatedActiveLoad": null,
//         "estimatedReactiveLoad": null,
//         "fullPower": null,
//         "i": 146.79640711660966
//     }, "drilled": { "n": 0.92, "cos": 0.9, "u": 0.38, "count": 2, "ph": 14, "np": 28, "kv": 0.12, "tg": 1, "npk": 3.36, "npktg": 3.36, "np2": 392, "efficiencyQuantity": null, "activPowerCoefficient": null, "estimatedActiveLoad": null, "estimatedReactiveLoad": null, "fullPower": null, "i": 51.37874249081339 }, "jointing": { "n": 0.92, "cos": 0.9, "u": 0.38, "count": 4, "ph": 42, "np": 168, "kv": 0.15, "tg": 1.33, "npk": 25.2, "npktg": 33.516, "np2": 7056, "efficiencyQuantity": null, "activPowerCoefficient": null, "estimatedActiveLoad": null, "estimatedReactiveLoad": null, "fullPower": null, "i": 308.2724549448803 }, "circularSaw": { "n": 0.92, "cos": 0.9, "u": 0.38, "count": 1, "ph": 36, "np": 36, "kv": 0.3, "tg": 1.52, "npk": 10.799999999999999, "npktg": 16.415999999999997, "np2": 1296, "efficiencyQuantity": null, "activPowerCoefficient": null, "estimatedActiveLoad": null, "estimatedReactiveLoad": null, "fullPower": null, "i": 66.05838320247435 }, "press": { "n": 0.92, "cos": 0.9, "u": 0.38, "count": 1, "ph": 20, "np": 20, "kv": 0.5, "tg": 0.75, "npk": 10, "npktg": 7.5, "np2": 400, "efficiencyQuantity": null, "activPowerCoefficient": null, "estimatedActiveLoad": null, "estimatedReactiveLoad": null, "fullPower": null, "i": 36.699101779152414 }, "polishing": { "n": 0.92, "cos": 0.9, "u": 0.38, "count": 1, "ph": 40, "np": 40, "kv": 0.2, "tg": 1, "npk": 8, "npktg": 8, "np2": 1600, "efficiencyQuantity": null, "activPowerCoefficient": null, "estimatedActiveLoad": null, "estimatedReactiveLoad": null, "fullPower": null, "i": 73.39820355830483 }, "shaper": { "n": 0.92, "cos": 0.9, "u": 0.38, "count": 2, "ph": 32, "np": 64, "kv": 0.2, "tg": 1, "npk": 12.8, "npktg": 12.8, "np2": 2048, "efficiencyQuantity": null, "activPowerCoefficient": null, "estimatedActiveLoad": null, "estimatedReactiveLoad": null, "fullPower": null, "i": 117.43712569328774 }, "fan": { "n": 0.92, "cos": 0.9, "u": 0.38, "count": 1, "ph": 20, "np": 20, "kv": 0.65, "tg": 0.75, "npk": 13, "npktg": 9.75, "np2": 400, "efficiencyQuantity": null, "activPowerCoefficient": null, "estimatedActiveLoad": null, "estimatedReactiveLoad": null, "fullPower": null, "i": 36.699101779152414 }, "shr1": { "n": null, "cos": null, "u": null, "count": 16, "ph": null, "np": 456, "kv": 0.20868421052631578, "tg": null, "npk": 95.16, "npktg": 107.30199999999999, "np2": 14792, "efficiencyQuantity": 14.057328285559763, "activPowerCoefficient": 1.25, "estimatedActiveLoad": 118.94999999999999, "estimatedReactiveLoad": 107.30199999999999, "fullPower": 160.19619753290024, "i": 313.02631578947364 }, "shr2": { "n": null, "cos": null, "u": null, "count": 16, "ph": null, "np": 456, "kv": 0.20868421052631578, "tg": null, "npk": 95.16, "npktg": 107.30199999999999, "np2": 14792, "efficiencyQuantity": 14.057328285559763, "activPowerCoefficient": 1.25, "estimatedActiveLoad": 118.94999999999999, "estimatedReactiveLoad": 107.30199999999999, "fullPower": 160.19619753290024, "i": 313.02631578947364 }, "shr3": { "n": null, "cos": null, "u": null, "count": 16, "ph": null, "np": 456, "kv": 0.20868421052631578, "tg": null, "npk": 95.16, "npktg": 107.30199999999999, "np2": 14792, "efficiencyQuantity": 14.057328285559763, "activPowerCoefficient": 1.25, "estimatedActiveLoad": 118.94999999999999, "estimatedReactiveLoad": 107.30199999999999, "fullPower": 160.19619753290024, "i": 313.02631578947364 }, "welding": { "n": 0.92, "cos": 0.9, "u": 0.38, "count": 2, "ph": 100, "np": 200, "kv": 0.2, "tg": 3, "npk": 40, "npktg": 120, "np2": 20000, "efficiencyQuantity": null, "activPowerCoefficient": null, "estimatedActiveLoad": null, "estimatedReactiveLoad": null, "fullPower": null, "i": 366.99101779152414 }, "dryer": { "n": 0.92, "cos": 0.9, "u": 0.38, "count": 2, "ph": 120, "np": 240, "kv": 0.8, "tg": 0, "npk": 192, "npktg": 0, "np2": 28800, "efficiencyQuantity": null, "activPowerCoefficient": null, "estimatedActiveLoad": null, "estimatedReactiveLoad": null, "fullPower": null, "i": 440.38922134982903 }, "fullLoad": { "n": null, "cos": null, "u": null, "count": 81, "ph": null, "np": 2330, "kv": 0.32274678111587984, "tg": null, "npk": 752, "npktg": 657, "np2": 96388, "efficiencyQuantity": 56.32340125326804, "activPowerCoefficient": 0.7, "estimatedActiveLoad": 526.4, "estimatedReactiveLoad": 459.9, "fullPower": 699.0028397653331, "i": 1385.2631578947369 }
// }

const toNumber = (val: unknown) =>
    typeof val === "string" && val.trim() !== "" ? parseFloat(val) : undefined;

const initialDevices = [
    "grinding", "drilled", "jointing", "circularSaw", "press", "polishing",
    "shaper", "fan", "welding", "dryer"
];

const initialValues = {
    n: 0.92,
    cos: 0.9,
    u: 0.38,
    count: "",
    ph: "",
    kv: "",
    tg: ""
};



export default function Task6() {
    const [task6Result, setTask6Result] = useState<Record<string, any> | null>(null);

    const [devices, setDevices] = useState(
        Object.fromEntries(initialDevices.map(device => [device, { ...initialValues }]))
    );

    const handleChange = (device, field, value) => {
        setDevices(prev => ({
            ...prev,
            [device]: {
                ...prev[device],
                [field]: value
            }
        }));
    };

    const generateRequest = () => {
        const formattedData = Object.fromEntries(
            Object.entries(devices).map(([device, values]) => [
                device,
                Object.fromEntries(
                    Object.entries(values).map(([key, val]) => [key, val !== "" ? parseFloat(val) : null])
                )
            ])
        );
        console.log("Generated JSON Request:", JSON.stringify(formattedData, null, 2));
        axios.post('http://localhost:8080/task6', JSON.stringify(formattedData, null, 2), {
            headers: { 'Content-Type': 'application/json' },
        })
            .then((response) => {
                console.log('Success:', response.data);
                setTask6Result(response.data);
            })
            .catch((error) => {
                console.error('Error:', error);
            });
    };



    return (
        <section className="flex flex-col px-10 gap-6 justify-center">
        <div>
                <h1 className="text-xl font-bold mb-4 text-center">Enter Device Data</h1>
                {Object.keys(devices).map(device => (
                    <div key={device} className="p-4 mb-4">
                        <h2 className="text-lg font-semibold mb-2">{device}</h2>
                        <div className="flex gap-2">
                            {Object.keys(initialValues).map(field => (
                                <div key={field} className="">
                                    <Input
                                        value={devices[device][field]}
                                        placeholder={field}
                                        onChange={(e) => handleChange(device, field, e.target.value)}
                                        className="border p-1 rounded"
                                    />
                                </div>
                            ))}
                        </div>
                    </div>
                ))}
                <Button
                    onClick={generateRequest}
                >
                    Calculate
                </Button>
            </div>
        <div className="w-full">
        <h1 className="text-xl font-bold mb-4 text-center">Results:</h1>
                {task6Result ? (
                    <div className="border rounded-lg overflow-hidden shadow-lg">
                        <table className="w-full border-collapse">
                            <thead>
                                <tr>
                                    <th className="border text-center">Electric Device</th>
                                    {Object.keys(task6Result[Object.keys(task6Result)[0]]).map((key) => (
                                        <th key={key} className="border p-1 text-center">
                                            <div className="transform whitespace-nowrap">{key.charAt(0).toUpperCase() + key.slice(1)}</div>
                                        </th>
                                    ))}
                                </tr>
                            </thead>
                            <tbody>
                                {Object.entries(task6Result).map(([device, values], index) => (
                                    <tr key={device}>
                                        <td className="border p-2">{device}</td>
                                        {Object.entries(values).map(([key, value], index) => (
                                            <td key={key} className="border p-2">{value === 0 || value === null ? "-" : Math.round(value * 10) / 10}</td>
                                        ))}
                                    </tr>
                                ))}
                            </tbody>
                        </table>
                    </div>
                ) : (
                    <p className="text-gray-500 text-center">No Data</p>
                )}
        </div>
        </section>
    );
}
