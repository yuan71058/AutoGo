#ifndef YOLO_H
#define YOLO_H

#ifdef __cplusplus
extern "C" {
#endif

typedef struct Yolo Yolo;

Yolo* newYolo();

const char *loadModelYolo(Yolo *obj, const char *version, const char *param_path, const char *bin_path, const char* labels, int num_threads);

char *detectYolo(Yolo *obj, const unsigned char *bitmapData, int width, int height, float nms, float prob, int size);

void closeYolo(Yolo *obj);

#ifdef __cplusplus
}
#endif

#endif // YOLO_H
