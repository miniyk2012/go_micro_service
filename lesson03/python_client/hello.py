from __future__ import print_function

import logging

from google.protobuf.json_format import MessageToJson

logging.basicConfig(format='%(asctime)s %(message)s', level=logging.INFO)

import grpc
from add_pb2 import AddRequest, AddResponse
import add_pb2_grpc


def run():
    # NOTE(gRPC Python Team): .close() is possible on a channel and should be
    # used in circumstances in which the with statement does not fit the needs
    # of the code.
    with grpc.insecure_channel('127.0.0.1:8972') as channel:
        stub = add_pb2_grpc.AddServiceStub(channel)
        resp: AddResponse = stub.Add(AddRequest(a=2, b=3))
    logging.info(MessageToJson(resp))

if __name__ == "__main__":
    run()
