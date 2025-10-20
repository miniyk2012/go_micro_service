import datetime

from google.protobuf import any_pb2 as _any_pb2
from google.protobuf import timestamp_pb2 as _timestamp_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from collections.abc import Iterable as _Iterable, Mapping as _Mapping
from typing import ClassVar as _ClassVar, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class Week(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    UNKNOWN: _ClassVar[Week]
    MONDAY: _ClassVar[Week]
UNKNOWN: Week
MONDAY: Week

class SearchRequest(_message.Message):
    __slots__ = ("query", "page_number", "num", "score", "corpus", "DateOfBirth")
    class Corpus(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
        __slots__ = ()
        UNIVERSAL: _ClassVar[SearchRequest.Corpus]
        WEB: _ClassVar[SearchRequest.Corpus]
        IMAGES: _ClassVar[SearchRequest.Corpus]
        LOCAL: _ClassVar[SearchRequest.Corpus]
        NEWS: _ClassVar[SearchRequest.Corpus]
        PRODUCTS: _ClassVar[SearchRequest.Corpus]
        VIDEO: _ClassVar[SearchRequest.Corpus]
    UNIVERSAL: SearchRequest.Corpus
    WEB: SearchRequest.Corpus
    IMAGES: SearchRequest.Corpus
    LOCAL: SearchRequest.Corpus
    NEWS: SearchRequest.Corpus
    PRODUCTS: SearchRequest.Corpus
    VIDEO: SearchRequest.Corpus
    QUERY_FIELD_NUMBER: _ClassVar[int]
    PAGE_NUMBER_FIELD_NUMBER: _ClassVar[int]
    NUM_FIELD_NUMBER: _ClassVar[int]
    SCORE_FIELD_NUMBER: _ClassVar[int]
    CORPUS_FIELD_NUMBER: _ClassVar[int]
    DATEOFBIRTH_FIELD_NUMBER: _ClassVar[int]
    query: str
    page_number: int
    num: int
    score: float
    corpus: SearchRequest.Corpus
    DateOfBirth: _timestamp_pb2.Timestamp
    def __init__(self, query: _Optional[str] = ..., page_number: _Optional[int] = ..., num: _Optional[int] = ..., score: _Optional[float] = ..., corpus: _Optional[_Union[SearchRequest.Corpus, str]] = ..., DateOfBirth: _Optional[_Union[datetime.datetime, _timestamp_pb2.Timestamp, _Mapping]] = ...) -> None: ...

class SearchResponse(_message.Message):
    __slots__ = ("ret", "corpus", "result", "map_field", "details")
    class MapFieldEntry(_message.Message):
        __slots__ = ("key", "value")
        KEY_FIELD_NUMBER: _ClassVar[int]
        VALUE_FIELD_NUMBER: _ClassVar[int]
        key: str
        value: int
        def __init__(self, key: _Optional[str] = ..., value: _Optional[int] = ...) -> None: ...
    RET_FIELD_NUMBER: _ClassVar[int]
    CORPUS_FIELD_NUMBER: _ClassVar[int]
    RESULT_FIELD_NUMBER: _ClassVar[int]
    MAP_FIELD_FIELD_NUMBER: _ClassVar[int]
    DETAILS_FIELD_NUMBER: _ClassVar[int]
    ret: _containers.RepeatedScalarFieldContainer[str]
    corpus: SearchRequest.Corpus
    result: Result
    map_field: _containers.ScalarMap[str, int]
    details: _containers.RepeatedCompositeFieldContainer[_any_pb2.Any]
    def __init__(self, ret: _Optional[_Iterable[str]] = ..., corpus: _Optional[_Union[SearchRequest.Corpus, str]] = ..., result: _Optional[_Union[Result, _Mapping]] = ..., map_field: _Optional[_Mapping[str, int]] = ..., details: _Optional[_Iterable[_Union[_any_pb2.Any, _Mapping]]] = ...) -> None: ...

class Result(_message.Message):
    __slots__ = ("url", "title", "snippets", "week")
    URL_FIELD_NUMBER: _ClassVar[int]
    TITLE_FIELD_NUMBER: _ClassVar[int]
    SNIPPETS_FIELD_NUMBER: _ClassVar[int]
    WEEK_FIELD_NUMBER: _ClassVar[int]
    url: str
    title: str
    snippets: _containers.RepeatedScalarFieldContainer[str]
    week: Week
    def __init__(self, url: _Optional[str] = ..., title: _Optional[str] = ..., snippets: _Optional[_Iterable[str]] = ..., week: _Optional[_Union[Week, str]] = ...) -> None: ...

class SampleMessage(_message.Message):
    __slots__ = ("name", "sub_message")
    class SubMessage(_message.Message):
        __slots__ = ("age",)
        AGE_FIELD_NUMBER: _ClassVar[int]
        age: int
        def __init__(self, age: _Optional[int] = ...) -> None: ...
    NAME_FIELD_NUMBER: _ClassVar[int]
    SUB_MESSAGE_FIELD_NUMBER: _ClassVar[int]
    name: str
    sub_message: SampleMessage.SubMessage
    def __init__(self, name: _Optional[str] = ..., sub_message: _Optional[_Union[SampleMessage.SubMessage, _Mapping]] = ...) -> None: ...
