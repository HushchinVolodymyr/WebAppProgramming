import Task1 from "@/components/blocks/task1/task1";
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs"

export default function Home() {
  return (
    <main className={"flex flex-col items-center h-screen"}>
      <Tabs defaultValue="task1" className="w-screen flex flex-col items-center">
        <TabsList className={`mx-auto`}>
          <TabsTrigger value="task1">Task 1</TabsTrigger>
          <TabsTrigger value="task2">Task 2</TabsTrigger>
          <TabsTrigger value="task3">Task 3</TabsTrigger>
          <TabsTrigger value="task4">Task 4</TabsTrigger>
          <TabsTrigger value="task5">Task 5</TabsTrigger>
          <TabsTrigger value="task6">Task 6</TabsTrigger>
        </TabsList>
        <TabsContent value="task1" className="w-full">
          <Task1 />
        </TabsContent>
        <TabsContent value="task2">
          <h1>Task 2</h1>
        </TabsContent>
      </Tabs>
    </main>
  );
}
