import Task1 from "@/components/blocks/task1/task1";
import Task2 from "@/components/blocks/task2/task2";
import Task3 from "@/components/blocks/task3/task3";
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs"

export default function Home() {
  return (
    <main className={"flex flex-col items-center h-screen"}>
      <Tabs defaultValue="task3" className="w-screen flex flex-col items-center">
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
        <TabsContent value="task2" className="w-full">
          <Task2 />
        </TabsContent>
        <TabsContent value="task3" className="w-full">
          <Task3 />
        </TabsContent>
      </Tabs>
    </main>
  );
}
