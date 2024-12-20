#include "extensions.h"
#include <taglib/fileref.h>
#include <taglib/mp4file.h>
#include <taglib/mp4tag.h>
#include <taglib/tstring.h>
#include <string>
#include <locale>
#include <codecvt>

static const char *BASE_MP4_ATOM = "----:com.apple.iTunes:";

void taglib_set_item_mp4(TagLib_File *file, const char *key, const char *value) {
    if(file == NULL || key == NULL || value == NULL)
        return;
    TagLib::MP4::File *mfile = dynamic_cast<TagLib::MP4::File *>(reinterpret_cast<TagLib::FileRef *>(file)->file());
    TagLib::MP4::Tag *tag = mfile->tag();
    if(tag) {
        TagLib::String tagKey(BASE_MP4_ATOM);
        tagKey.append(TagLib::String(key));
        tag->setItem(tagKey, TagLib::StringList(value));
    }
}

void taglib_set_picture(TagLib_File *file, const char *data, unsigned int size, const char *desc, const char *mime, const char *typ) {
    TAGLIB_COMPLEX_PROPERTY_PICTURE(props, data, size, desc, mime, typ);
    taglib_complex_property_set(file, "PICTURE", props);
}