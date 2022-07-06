import { ProGallery } from 'pro-gallery';
import { placeholder } from './utils/constants';

function App() {
  return (
    <div className="App container mx-auto">
      <div className="text-2xl text-white my-4">Choose data fetching method</div>
      <div className="flex-row justify-between p-5">
        <div className="flex-col"><button>REST</button></div>
        <div className="flex-col"><button>GraphQL</button></div>
        <div className="flex-col"><button>gRPC</button></div>
      </div>
      <div className="text-2xl text-white my-3">Stats</div>
      <div className="stats flex-row">
        <div className="flex-col">
          <div className="flex-row text-xl text-white">Data fetching time: </div>
          <div className="flex-row text-xl text-white">Compress time: </div>
          <div className="flex-row"></div>
        </div>
      </div>
      <div className="text-2xl text-white my-3">Gallery</div>
      <div className="flex-row">
        <ProGallery
          items={[  { // Image Item:
            itemId: '1',
            mediaUrl: placeholder,
            metaData: {
              type: 'image',
              height: 200,
              width: 100,
              title: 'sample-title',
              description: 'sample-description',
              focalPoint: [0, 0],
              link: {
                url: placeholder,
                target: '_blank'
              },
            }
          },
          { // Video Item:
            itemId: '2',
            mediaUrl: placeholder,
            metaData: {
              type: 'image',
              height: 200,
              width: 100,
              poster: 'sample-image-url',
              title: 'sample-title',
              description: 'sample-description',
              focalPoint: [0, 0],
              link: {
                url: placeholder,
                target: '_blank'
              },
            }
          }]}
          options={{}}
          container={{
            width: 1000,
            height: 1500,
          }}
          scrollingElement={document.getElementById('gallery') || window}
          eventsListener={(eName, eData) => console.log({ eName, eData })}
        />
      </div>
    </div>
  );
}

export default App;
