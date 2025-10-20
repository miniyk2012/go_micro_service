from __future__ import print_function
import logging

import grpc
import demo_pb2
import demo_pb2_grpc


def run():
    # NOTE(gRPC Python Team): .close() is possible on a channel and should be
    # used in circumstances in which the with statement does not fit the needs
    # of the code.
    with grpc.insecure_channel('127.0.0.1:50051') as channel:
        stub = demo_pb2_grpc.SearchServiceStub(channel)
        resp: demo_pb2.SearchResponse = stub.Search(demo_pb2.SearchRequest(query='yangkai'))
    print("Greeter client received: " + str(resp.result))


if __name__ == '__main__':
    logging.basicConfig()
    run()
