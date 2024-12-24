from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Optional as _Optional

DESCRIPTOR: _descriptor.FileDescriptor

class AddRequest(_message.Message):
    __slots__ = ("a", "b")
    A_FIELD_NUMBER: _ClassVar[int]
    B_FIELD_NUMBER: _ClassVar[int]
    a: int
    b: int
    def __init__(self, a: _Optional[int] = ..., b: _Optional[int] = ...) -> None: ...

class AddResponse(_message.Message):
    __slots__ = ("result", "code", "msg")
    RESULT_FIELD_NUMBER: _ClassVar[int]
    CODE_FIELD_NUMBER: _ClassVar[int]
    MSG_FIELD_NUMBER: _ClassVar[int]
    result: int
    code: int
    msg: str
    def __init__(self, result: _Optional[int] = ..., code: _Optional[int] = ..., msg: _Optional[str] = ...) -> None: ...
