/**
 * v0 by Vercel.
 * @see https://v0.dev/t/PgjBM0G63Wj
 * Documentation: https://v0.dev/docs#integrating-generated-code-into-your-nextjs-app
 */
import { Card } from "@/components/ui/card"

export default function Component() {
  return (
    <div className="w-screen h-screen bg-gray-900 flex justify-center items-center">
      <div className="grid grid-cols-3 gap-4 w-full h-full p-4">
        <Card className="flex flex-col justify-between p-4 bg-green-200 h-full">
          <div className="text-lg font-bold">Self</div>
          <div className="text-3xl font-bold flex justify-end">0ms</div>
          <HeartPulseIcon className="w-6 h-6 text-gray-700" />
        </Card>
        <Card className="flex flex-col justify-between p-4 bg-green-200 h-full">
          <div className="text-lg font-bold">Self</div>
          <div className="flex justify-end">
            <ServerIcon className="w-6 h-6 text-gray-700" />
          </div>
        </Card>
        <Card className="flex flex-col justify-between p-4 bg-green-200 h-full">
          <div className="text-lg font-bold">Office1</div>
          <div className="text-3xl font-bold flex justify-end">317ms</div>
          <HeartPulseIcon className="w-6 h-6 text-gray-700" />
        </Card>
        <Card className="flex flex-col justify-between p-4 bg-green-200 h-full">
          <div className="text-lg font-bold">Server</div>
          <div className="text-3xl font-bold flex justify-end">68ms</div>
          <HeartPulseIcon className="w-6 h-6 text-gray-700" />
        </Card>
        <Card className="flex flex-col justify-between p-4 bg-green-200 h-full">
          <div className="text-lg font-bold">Server Responding</div>
          <div className="text-sm flex justify-end">HTTP</div>
        </Card>
        <Card className="flex flex-col justify-between p-4 bg-green-200 h-full">
          <div className="text-lg font-bold">Office2</div>
          <div className="text-3xl font-bold flex justify-end">334ms</div>
          <HeartPulseIcon className="w-6 h-6 text-gray-700" />
        </Card>
        <Card className="flex flex-col justify-between p-4 bg-green-200 h-full">
          <div className="text-lg font-bold">Office Synology</div>
          <div className="text-3xl font-bold flex justify-end">66ms</div>
          <HeartPulseIcon className="w-6 h-6 text-gray-700" />
        </Card>
        <Card className="flex flex-col justify-between p-4 bg-green-200 h-full">
          <div className="text-lg font-bold">Office Synology Responding</div>
          <div className="text-sm flex justify-end">HTTP</div>
        </Card>
        <Card className="flex flex-col justify-between p-4 bg-red-400 h-full">
          <div className="text-lg font-bold">Office3</div>
          <div className="flex justify-end">
            <HeartPulseIcon className="w-6 h-6 text-gray-700" />
          </div>
        </Card>
        <Card className="flex flex-col justify-between p-4 bg-green-200 h-full">
          <div className="text-lg font-bold">Home Synology</div>
          <div className="text-3xl font-bold flex justify-end">41ms</div>
          <HeartPulseIcon className="w-6 h-6 text-gray-700" />
        </Card>
        <Card className="flex flex-col justify-between p-4 bg-green-200 h-full">
          <div className="text-lg font-bold">Home Synology Responding</div>
          <div className="text-sm flex justify-end">HTTP</div>
        </Card>
      </div>
    </div>
  )
}

function HeartPulseIcon(props) {
  return (
    <svg
      {...props}
      xmlns="http://www.w3.org/2000/svg"
      width="24"
      height="24"
      viewBox="0 0 24 24"
      fill="none"
      stroke="currentColor"
      strokeWidth="2"
      strokeLinecap="round"
      strokeLinejoin="round"
    >
      <path d="M19 14c1.49-1.46 3-3.21 3-5.5A5.5 5.5 0 0 0 16.5 3c-1.76 0-3 .5-4.5 2-1.5-1.5-2.74-2-4.5-2A5.5 5.5 0 0 0 2 8.5c0 2.3 1.5 4.05 3 5.5l7 7Z" />
      <path d="M3.22 12H9.5l.5-1 2 4.5 2-7 1.5 3.5h5.27" />
    </svg>
  )
}


function ServerIcon(props) {
  return (
    <svg
      {...props}
      xmlns="http://www.w3.org/2000/svg"
      width="24"
      height="24"
      viewBox="0 0 24 24"
      fill="none"
      stroke="currentColor"
      strokeWidth="2"
      strokeLinecap="round"
      strokeLinejoin="round"
    >
      <rect width="20" height="8" x="2" y="2" rx="2" ry="2" />
      <rect width="20" height="8" x="2" y="14" rx="2" ry="2" />
      <line x1="6" x2="6.01" y1="6" y2="6" />
      <line x1="6" x2="6.01" y1="18" y2="18" />
    </svg>
  )
}
