"use client";

import { zodResolver } from "@hookform/resolvers/zod";
import { useForm } from "react-hook-form";
import { z } from "zod";
import {
    Form,
    FormControl,
    FormDescription,
    FormField,
    FormItem,
    FormLabel,
    FormMessage,
} from "@/components/ui/form";
import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";
import axios from 'axios';
import { useState } from "react";

const toNumber = (val: unknown) =>
    typeof val === "string" && val.trim() !== "" ? parseFloat(val) : undefined;

// Schema task 1 part 1 for the form
const task1Part1FormSchema = z.object({
    hydrogen: z.preprocess(toNumber, z.number().min(0, "Must be a number")),
    carbon: z.preprocess(toNumber, z.number().min(0, "Must be a number")),
    sulfur: z.preprocess(toNumber, z.number().min(0, "Must be a number")),
    nitrogen: z.preprocess(toNumber, z.number().min(0, "Must be a number")),
    oxygen: z.preprocess(toNumber, z.number().min(0, "Must be a number")),
    moister: z.preprocess(toNumber, z.number().min(0, "Must be a number")),
    ash: z.preprocess(toNumber, z.number().min(0, "Must be a number")),
});

// Schema task 1 part 2 for the form
const task1Part2FormSchema = z.object({
    carbon: z.preprocess(toNumber, z.number().min(0, "Must be a number")),
    hydrogen: z.preprocess(toNumber, z.number().min(0, "Must be a number")),
    oxygen: z.preprocess(toNumber, z.number().min(0, "Must be a number")),
    sulfur: z.preprocess(toNumber, z.number().min(0, "Must be a number")),
    lowerHeatingValue: z.preprocess(toNumber, z.number().min(0, "Must be a number")),
    moister: z.preprocess(toNumber, z.number().min(0, "Must be a number")),
    ash: z.preprocess(toNumber, z.number().min(0, "Must be a number")),
    vanadium: z.preprocess(toNumber, z.number().min(0, "Must be a number")),
});


export default function Task1() {
    // Result state for task 1 part 1
    const [task1Part1Result, setTask1Part1Result] = useState<Record<string, any> | null>(null);
    // Result state for task 1 part 2
    const [task1Part2Result, setTask1Part2Result] = useState<Record<string, any> | null>(null);

    // Form instance for task 1 part 1
    const task1Part1Form = useForm({
        resolver: zodResolver(task1Part1FormSchema),
    });

    // Form instance for task 1 part 2
    const task1Part2Form = useForm({
        resolver: zodResolver(task1Part2FormSchema),
    });


    // Submit task 1 part 1
    function onSubmitTask1Part1(data: z.infer<typeof task1Part1FormSchema>) {
        // Send data to the server
        axios.post('http://localhost:8080/task1/part1', data, {
            headers: {
                'Content-Type': 'application/json',
            },
        })
            // Handle response
            .then((response) => {
                console.log('Success:', response.data);
                // Set the result
                setTask1Part1Result(response.data);
            })
            // Handle error
            .catch((error) => {
                console.error('Error:', error);
            });
    }

    // Submit task 1 part 2
    function onSubmitTask1Part2(data: z.infer<typeof task1Part2FormSchema>) {
        // Send data to the server
        axios.post('http://localhost:8080/task1/part2', data, {
            headers: {
                'Content-Type': 'application/json',
            },
        })
            // Handle response
            .then((response) => {
                console.log('Success:', response.data);
                // Set the result
                setTask1Part2Result(response.data);
            })
            // Handle error
            .catch((error) => {
                console.error('Error:', error);
            });
    }

    return (
        <div className={`flex w-full`}>
            <section className={`w-1/2 flex px-10 gap-6`}>
                <div>
                    <Form {...task1Part1Form}>
                        <form onSubmit={task1Part1Form.handleSubmit(onSubmitTask1Part1)} className="flex flex-col items-center w-full">
                            {["hydrogen", "carbon", "sulfur", "nitrogen", "oxygen", "moister", "ash"].map((field) => (
                                <FormField
                                    key={field}
                                    control={task1Part1Form.control}
                                    name={field as keyof typeof task1Part1FormSchema.shape}
                                    render={({ field }) => (
                                        <FormItem>
                                            <FormLabel>{field.name.charAt(0).toUpperCase() + field.name.slice(1)}</FormLabel>
                                            <FormControl>
                                                <Input
                                                    type="number"
                                                    placeholder={field.name}
                                                    {...field}
                                                    value={field.value !== undefined ? field.value : ""}
                                                    onChange={(e) => field.onChange(e.target.value)}
                                                    style={{ width: "300px" }}
                                                />
                                            </FormControl>
                                            <FormMessage />
                                        </FormItem>
                                    )}
                                />
                            ))}

                            <Button
                                type="submit"
                                className={`mt-2`}
                            >
                                Calculate Task 1
                            </Button>
                        </form>
                    </Form>
                </div>

                <div>
                    <h1 className={`text-xl`}>Results:</h1>
                    {task1Part1Result ? (
                        <div className="border p-4 rounded ">
                            {Object.entries(task1Part1Result).map(([key, value]) => (
                                <div key={key} className="flex justify-between py-1">
                                    <span className="font-medium">{key}:</span>
                                    <span>{value}</span>
                                </div>
                            ))}
                        </div>
                    ) : null}

                </div>
            </section>

            <section className={`w-1/2 flex px-10 gap-6`}>
                <div>
                    <Form {...task1Part2Form}>
                        <form onSubmit={task1Part2Form.handleSubmit(onSubmitTask1Part2)} className="flex flex-col items-center w-full">
                            {["carbon", "hydrogen", "oxygen", "sulfur", "lowerHeatingValue", "moister", "ash", "vanadium"].map((field) => (
                                <FormField
                                    key={field}
                                    control={task1Part2Form.control}
                                    name={field as keyof typeof task1Part2FormSchema.shape}
                                    render={({ field }) => (
                                        <FormItem>
                                            <FormLabel>{field.name.charAt(0).toUpperCase() + field.name.slice(1)}</FormLabel>
                                            <FormControl>
                                                <Input
                                                    type="number"
                                                    placeholder={field.name}
                                                    {...field}
                                                    value={field.value !== undefined ? field.value : ""}
                                                    onChange={(e) => field.onChange(e.target.value)}
                                                    style={{ width: "300px" }}
                                                />
                                            </FormControl>
                                            <FormMessage />
                                        </FormItem>
                                    )}
                                />
                            ))}
                            <Button
                                type="submit"
                                className={`mt-2`}
                            >
                                Calculate Task 2
                            </Button>
                        </form>
                    </Form>
                </div>
                <div>
                    <h1 className={`text-xl`}>Results:</h1>
                    {task1Part2Result ? (
                        <div className="border p-4 rounded ">
                            {Object.entries(task1Part2Result).map(([key, value]) => (
                                <div key={key} className="flex justify-between py-1">
                                    <span className="font-medium">{key}:</span>
                                    <span>{value}</span>
                                </div>
                            ))}
                        </div>
                    ) : null}

                </div>

            </section>
        </div>
    );
}


