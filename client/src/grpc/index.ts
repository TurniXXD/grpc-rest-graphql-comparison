require('dotenv').config()
const grpc = require('grpc')
const protoLoader = require('@grpc/proto-loader')
const packageDef = protoLoader.loadSync('proto/gallery.proto', {})
const grpcObject = grpc.loadPackageDefinition(packageDef)
const gallery = grpcObject.gallery

const client = new gallery.Gallery(process.env.NODE_API, grpc.credentials.createInsecure())

// -1 => creating a new ID but a new item
export const addImage = (data: any) => {
	client.addImage(
		{
			id: -1,
			url: data.url || '',
			width: data.width || 0,
			height: data.height || 0,
			alt: data.alt,
		},
		(err: Error, res: any) => {
			err ? console.log('err' + JSON.stringify(err)) : console.log('res:' + JSON.stringify(res))
		}
	)
}
